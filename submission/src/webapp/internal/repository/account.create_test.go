package repository

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestAccount_CreateAccount(t *testing.T) {
	type args struct {
		email    string
		password string
	}

	testCases := []struct {
		name          string
		args          args
		mock          func(mock sqlmock.Sqlmock)
		expectedError string
	}{
		{
			name: "successful account creation",
			args: args{
				email:    "test@example.com",
				password: "securepassword",
			},
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(regexp.QuoteMeta(accountCreationQuery)).
					WithArgs("test@example.com", "securepassword").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: "",
		},
		{
			name: "duplicate email",
			args: args{
				email:    "exists@example.com",
				password: "securepassword",
			},
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(regexp.QuoteMeta(accountCreationQuery)).
					WithArgs("exists@example.com", "securepassword").
					WillReturnError(errors.New("pq: duplicate key value violates unique constraint"))
			},
			expectedError: "failed to insert account: pq: duplicate key value violates unique constraint",
		},
		{
			name: "no rows affected",
			args: args{
				email:    "test@example.com",
				password: "securepassword",
			},
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(regexp.QuoteMeta(accountCreationQuery)).
					WithArgs("test@example.com", "securepassword").
					WillReturnResult(sqlmock.NewResult(0, 0))
			},
			expectedError: "failed to insert account: no rows affected",
		},
		{
			name: "database connection error",
			args: args{
				email:    "test@example.com",
				password: "securepassword",
			},
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(regexp.QuoteMeta(accountCreationQuery)).
					WithArgs("test@example.com", "securepassword").
					WillReturnError(errors.New("connection failed"))
			},
			expectedError: "failed to insert account: connection failed",
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
			err = repo.CreateAccount(context.Background(), tc.args.email, tc.args.password)

			if tc.expectedError != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tc.expectedError)
			} else {
				assert.NoError(t, err)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
