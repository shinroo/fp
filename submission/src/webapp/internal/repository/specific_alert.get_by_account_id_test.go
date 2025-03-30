package repository

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/shinroo/fp/src/webapp/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestSpecificAlert_GetByAccountID(t *testing.T) {
	testAlerts := []*models.SpecificAlert{
		{
			ID:        1,
			AccountID: 123,
			Breed:     "Labrador",
			Gender:    models.Male,
			LifeStage: models.Adult,
		},
		{
			ID:        2,
			AccountID: 123,
			Breed:     "Poodle",
			Gender:    models.Female,
			LifeStage: models.Senior,
		},
	}

	testCases := []struct {
		name           string
		accountID      int
		mock           func(mock sqlmock.Sqlmock)
		expectedResult []*models.SpecificAlert
		expectedError  string
	}{
		{
			name:      "successful retrieval with alerts",
			accountID: 123,
			mock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "account_id", "breed", "gender", "life_stage"}).
					AddRow(testAlerts[0].ID, testAlerts[0].AccountID, testAlerts[0].Breed, testAlerts[0].Gender, testAlerts[0].LifeStage).
					AddRow(testAlerts[1].ID, testAlerts[1].AccountID, testAlerts[1].Breed, testAlerts[1].Gender, testAlerts[1].LifeStage)

				mock.ExpectQuery(regexp.QuoteMeta(specificAlertGetByAccountIDQuery)).
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
				rows := sqlmock.NewRows([]string{"id", "account_id", "breed", "gender", "life_stage"})
				mock.ExpectQuery(regexp.QuoteMeta(specificAlertGetByAccountIDQuery)).
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
				mock.ExpectQuery(regexp.QuoteMeta(specificAlertGetByAccountIDQuery)).
					WithArgs(789).
					WillReturnError(errors.New("connection failed"))
			},
			expectedResult: nil,
			expectedError:  "failed to get specific alerts by account id: connection failed",
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
			repo := NewSpecificAlert(sqlxDB)

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
