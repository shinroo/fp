package repository

import (
	"context"
	"database/sql"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/shinroo/fp/src/webapp/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestAccount_GetAccountByEmail(t *testing.T) {
	now := time.Now()
	testAccount := &models.Account{
		ID:        1,
		Email:     "test@example.com",
		Password:  "hashedpassword",
		CreatedAt: now,
		UpdatedAt: now,
	}

	type args struct {
		email string
	}

	testCases := []struct {
		name           string
		args           args
		mock           func(mock sqlmock.Sqlmock)
		expectedResult *models.Account
		expectedError  string
	}{
		{
			name: "successful account retrieval",
			args: args{
				email: "test@example.com",
			},
			mock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "email", "password", "created_at", "updated_at"}).
					AddRow(testAccount.ID, testAccount.Email, testAccount.Password, testAccount.CreatedAt, testAccount.UpdatedAt)

				mock.ExpectQuery(regexp.QuoteMeta(accountGetByEmailQuery)).
					WithArgs("test@example.com").
					WillReturnRows(rows)
			},
			expectedResult: testAccount,
			expectedError:  "",
		},
		{
			name: "account not found",
			args: args{
				email: "notfound@example.com",
			},
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(accountGetByEmailQuery)).
					WithArgs("notfound@example.com").
					WillReturnError(sql.ErrNoRows)
			},
			expectedResult: nil,
			expectedError:  "failed to get account by email: sql: no rows in result set",
		},
		{
			name: "database error",
			args: args{
				email: "test@example.com",
			},
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(accountGetByEmailQuery)).
					WithArgs("test@example.com").
					WillReturnError(errors.New("connection failed"))
			},
			expectedResult: nil,
			expectedError:  "failed to get account by email: connection failed",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			sqlxDB := sqlx.NewDb(db, "sqlmock")
			tc.mock(mock)

			repo := NewAccount(sqlxDB)
			result, err := repo.GetAccountByEmail(context.Background(), tc.args.email)

			if tc.expectedError != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tc.expectedError)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tc.expectedResult, result)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
