package repository

import (
	"context"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/shinroo/fp/src/webapp/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestSimilarityAlert_GetByAccountID(t *testing.T) {
	now := time.Now()
	testAlerts := []*models.SimilarityAlert{
		{
			ID:                  1,
			AccountID:           123,
			SimilarityThreshold: 0.85,
			Actioned:            false,
			CreatedAt:           now,
		},
		{
			ID:                  2,
			AccountID:           123,
			SimilarityThreshold: 0.90,
			Actioned:            true,
			CreatedAt:           now.Add(-24 * time.Hour),
		},
	}

	testCases := []struct {
		name           string
		accountID      int
		mock           func(mock sqlmock.Sqlmock)
		expectedResult []*models.SimilarityAlert
		expectedError  string
	}{
		{
			name:      "successful retrieval with multiple alerts",
			accountID: 123,
			mock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "account_id", "similarity_threshold", "actioned", "created_at"}).
					AddRow(testAlerts[0].ID, testAlerts[0].AccountID, testAlerts[0].SimilarityThreshold, testAlerts[0].Actioned, testAlerts[0].CreatedAt).
					AddRow(testAlerts[1].ID, testAlerts[1].AccountID, testAlerts[1].SimilarityThreshold, testAlerts[1].Actioned, testAlerts[1].CreatedAt)

				mock.ExpectQuery(regexp.QuoteMeta(similarityAlertGetByAccountIDQuery)).
					WithArgs(123).
					WillReturnRows(rows)
			},
			expectedResult: testAlerts,
			expectedError:  "",
		},
		{
			name:      "successful retrieval with no alerts",
			accountID: 456,
			mock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "account_id", "similarity_threshold", "actioned", "created_at"})
				mock.ExpectQuery(regexp.QuoteMeta(similarityAlertGetByAccountIDQuery)).
					WithArgs(456).
					WillReturnRows(rows)
			},
			expectedResult: nil,
			expectedError:  "",
		},
		{
			name:      "database error",
			accountID: 789,
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(similarityAlertGetByAccountIDQuery)).
					WithArgs(789).
					WillReturnError(errors.New("connection failed"))
			},
			expectedResult: nil,
			expectedError:  "failed to get similarity alerts by account id: connection failed",
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

			result, err := repo.GetByAccountID(context.Background(), tc.accountID)

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
