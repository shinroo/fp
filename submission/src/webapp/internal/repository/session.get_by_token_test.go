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

func TestSession_GetSessionByToken(t *testing.T) {
	now := time.Now()
	testSession := &models.Session{
		Token:     "test-token-123",
		AccountID: 123,
		CreatedAt: now,
	}

	testCases := []struct {
		name           string
		token          string
		mock           func(mock sqlmock.Sqlmock)
		expectedResult *models.Session
		expectedError  string
	}{
		{
			name:  "successful session retrieval",
			token: "test-token-123",
			mock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"token", "account_id", "created_at"}).
					AddRow(testSession.Token, testSession.AccountID, testSession.CreatedAt)
				mock.ExpectQuery(regexp.QuoteMeta(sessionGetByToken)).
					WithArgs("test-token-123").
					WillReturnRows(rows)
			},
			expectedResult: testSession,
			expectedError:  "",
		},
		{
			name:  "session not found",
			token: "non-existent-token",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(sessionGetByToken)).
					WithArgs("non-existent-token").
					WillReturnError(sql.ErrNoRows)
			},
			expectedResult: nil,
			expectedError:  "failed to get session by token: sql: no rows in result set",
		},
		{
			name:  "database error",
			token: "test-token-123",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(sessionGetByToken)).
					WithArgs("test-token-123").
					WillReturnError(errors.New("connection failed"))
			},
			expectedResult: nil,
			expectedError:  "failed to get session by token: connection failed",
		},
		{
			name:           "empty token",
			token:          "",
			mock:           func(_ sqlmock.Sqlmock) {},
			expectedResult: nil,
			expectedError:  "failed to get session by token: empty token",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("failed to create mock db: %v", err)
			}
			defer db.Close()

			sqlxDB := sqlx.NewDb(db, "sqlmock")
			repo := NewSession(sqlxDB)

			tc.mock(mock)

			result, err := repo.GetSessionByToken(context.Background(), tc.token)

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
