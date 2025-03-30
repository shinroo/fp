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

func TestSpecificAlert_CreateSpecificAlert(t *testing.T) {
	testCases := []struct {
		name          string
		accountID     int
		breed         string
		gender        models.Gender
		lifeStage     models.LifeStage
		mock          func(mock sqlmock.Sqlmock)
		expectedError string
	}{
		{
			name:      "successful alert creation",
			accountID: 123,
			breed:     "Labrador",
			gender:    models.Male,
			lifeStage: models.Adult,
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(regexp.QuoteMeta(specificAlertCreationQuery)).
					WithArgs(123, "Labrador", models.Adult, models.Male).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: "",
		},
		{
			name:      "duplicate alert",
			accountID: 456,
			breed:     "Beagle",
			gender:    models.Female,
			lifeStage: models.Puppy,
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(regexp.QuoteMeta(specificAlertCreationQuery)).
					WithArgs(456, "Beagle", models.Puppy, models.Female).
					WillReturnError(errors.New("duplicate key value violates unique constraint"))
			},
			expectedError: "failed to insert specific alert: duplicate key value violates unique constraint",
		},
		{
			name:      "no rows affected",
			accountID: 789,
			breed:     "Poodle",
			gender:    models.Male,
			lifeStage: models.Senior,
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(regexp.QuoteMeta(specificAlertCreationQuery)).
					WithArgs(789, "Poodle", models.Senior, models.Male).
					WillReturnResult(sqlmock.NewResult(0, 0))
			},
			expectedError: "failed to insert specific alert: no rows affected",
		},
		{
			name:      "database error",
			accountID: 999,
			breed:     "Bulldog",
			gender:    models.Male,
			lifeStage: models.Senior,
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(regexp.QuoteMeta(specificAlertCreationQuery)).
					WithArgs(999, "Bulldog", models.Senior, models.Male).
					WillReturnError(errors.New("connection failed"))
			},
			expectedError: "failed to insert specific alert: connection failed",
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

			err = repo.CreateSpecificAlert(context.Background(), tc.accountID, tc.breed, tc.gender, tc.lifeStage)

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
