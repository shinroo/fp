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

func TestSession_CreateSession(t *testing.T) {
	testCases := []struct {
		name          string
		accountID     int
		mock          func(mock sqlmock.Sqlmock)
		expectError   bool
		expectedError string
	}{
		{
			name:      "successful session creation",
			accountID: 123,
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(regexp.QuoteMeta(sessionCreateQuery)).
					WithArgs(sqlmock.AnyArg(), 123).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectError: false,
		},
		{
			name:      "duplicate token",
			accountID: 456,
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(regexp.QuoteMeta(sessionCreateQuery)).
					WithArgs(sqlmock.AnyArg(), 456).
					WillReturnError(errors.New("ERROR: duplicate key value violates unique constraint (SQLSTATE 23505)"))
			},
			expectError:   true,
			expectedError: "failed to insert session: ERROR: duplicate key value violates unique constraint (SQLSTATE 23505)",
		},
		{
			name:      "no rows affected",
			accountID: 789,
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(regexp.QuoteMeta(sessionCreateQuery)).
					WithArgs(sqlmock.AnyArg(), 789).
					WillReturnResult(sqlmock.NewResult(0, 0))
			},
			expectError:   true,
			expectedError: "failed to insert session: no rows affected",
		},
		{
			name:      "database connection error",
			accountID: 999,
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(regexp.QuoteMeta(sessionCreateQuery)).
					WithArgs(sqlmock.AnyArg(), 999).
					WillReturnError(errors.New("connection failed"))
			},
			expectError:   true,
			expectedError: "failed to insert session: connection failed",
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

			token, err := repo.CreateSession(context.Background(), tc.accountID)

			if tc.expectError {
				assert.Error(t, err)
				assert.EqualError(t, err, tc.expectedError)
				assert.Empty(t, token)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, token)
				assert.Greater(t, len(token), 30) // UUIDs are 36 chars long
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
