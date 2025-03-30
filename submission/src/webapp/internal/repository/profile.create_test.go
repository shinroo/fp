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

func TestProfile_CreateProfile(t *testing.T) {
	testCases := []struct {
		name          string
		accountID     int
		nearestSPCAID string
		mock          func(mock sqlmock.Sqlmock)
		expectedError string
	}{
		{
			name:          "successful profile creation",
			accountID:     123,
			nearestSPCAID: "SPCA123",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(regexp.QuoteMeta(profileCreationQuery)).
					WithArgs(123, "SPCA123").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: "",
		},
		{
			name:          "duplicate account ID",
			accountID:     123,
			nearestSPCAID: "SPCA123",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(regexp.QuoteMeta(profileCreationQuery)).
					WithArgs(123, "SPCA123").
					WillReturnError(errors.New("ERROR: duplicate key value violates unique constraint (SQLSTATE 23505)"))
			},
			expectedError: "failed to insert profile: ERROR: duplicate key value violates unique constraint (SQLSTATE 23505)",
		},
		{
			name:          "no rows affected",
			accountID:     456,
			nearestSPCAID: "SPCA456",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(regexp.QuoteMeta(profileCreationQuery)).
					WithArgs(456, "SPCA456").
					WillReturnResult(sqlmock.NewResult(0, 0))
			},
			expectedError: "failed to insert profile: no rows affected",
		},
		{
			name:          "database connection error",
			accountID:     789,
			nearestSPCAID: "SPCA789",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(regexp.QuoteMeta(profileCreationQuery)).
					WithArgs(789, "SPCA789").
					WillReturnError(errors.New("connection failed"))
			},
			expectedError: "failed to insert profile: connection failed",
		},
		{
			name:          "invalid SPCA ID format",
			accountID:     999,
			nearestSPCAID: "INVALID_ID",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(regexp.QuoteMeta(profileCreationQuery)).
					WithArgs(999, "INVALID_ID").
					WillReturnError(errors.New("invalid spca id format"))
			},
			expectedError: "failed to insert profile: invalid spca id format",
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
			repo := NewProfile(sqlxDB)

			tc.mock(mock)

			err = repo.CreateProfile(context.Background(), tc.accountID, tc.nearestSPCAID)

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
