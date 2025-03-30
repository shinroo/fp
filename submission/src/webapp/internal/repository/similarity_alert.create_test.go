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

func TestSimilarityAlert_CreateSimilarityAlert(t *testing.T) {
	testCases := []struct {
		name                string
		accountID           int
		similarityThreshold float32
		mock                func(mock sqlmock.Sqlmock)
		expectedError       string
	}{
		{
			name:                "successful alert creation",
			accountID:           123,
			similarityThreshold: 0.85,
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(regexp.QuoteMeta(similarityAlertCreationQuery)).
					WithArgs(123, float32(0.85)).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: "",
		},
		{
			name:                "duplicate alert",
			accountID:           456,
			similarityThreshold: 0.90,
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(regexp.QuoteMeta(similarityAlertCreationQuery)).
					WithArgs(456, float32(0.90)).
					WillReturnError(errors.New("duplicate key value violates unique constraint"))
			},
			expectedError: "failed to insert similarity alert: duplicate key value violates unique constraint",
		},
		{
			name:                "no rows affected",
			accountID:           789,
			similarityThreshold: 0.75,
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(regexp.QuoteMeta(similarityAlertCreationQuery)).
					WithArgs(789, float32(0.75)).
					WillReturnResult(sqlmock.NewResult(0, 0))
			},
			expectedError: "failed to insert similarity alert: no rows affected",
		},
		{
			name:                "database error",
			accountID:           999,
			similarityThreshold: 0.80,
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(regexp.QuoteMeta(similarityAlertCreationQuery)).
					WithArgs(999, float32(0.80)).
					WillReturnError(errors.New("connection failed"))
			},
			expectedError: "failed to insert similarity alert: connection failed",
		},
		{
			name:                "invalid threshold (too low)",
			accountID:           111,
			similarityThreshold: -0.1,
			mock:                func(_ sqlmock.Sqlmock) {},
			expectedError:       "failed to insert similarity alert: invalid similarity threshold",
		},
		{
			name:                "invalid threshold (too high)",
			accountID:           222,
			similarityThreshold: 1.1,
			mock:                func(_ sqlmock.Sqlmock) {},
			expectedError:       "failed to insert similarity alert: invalid similarity threshold",
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
			repo := NewSimilarityAlert(sqlxDB)

			tc.mock(mock)

			err = repo.CreateSimilarityAlert(context.Background(), tc.accountID, tc.similarityThreshold)

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
