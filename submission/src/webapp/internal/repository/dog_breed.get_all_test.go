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

func TestDogBreed_GetAll(t *testing.T) {
	testDogBreeds := []*models.DogBreed{
		{Name: "Labrador Retriever"},
		{Name: "Beagle"},
		{Name: "Poodle"},
	}

	testCases := []struct {
		name           string
		mock           func(mock sqlmock.Sqlmock)
		expectedResult []*models.DogBreed
		expectedError  string
	}{
		{
			name: "successful retrieval of all dog breeds",
			mock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"}).
					AddRow("Labrador Retriever").
					AddRow("Beagle").
					AddRow("Poodle")

				mock.ExpectQuery(regexp.QuoteMeta(dogBreedGetAllQuery)).
					WillReturnRows(rows)
			},
			expectedResult: testDogBreeds,
			expectedError:  "",
		},
		{
			name: "no dog breeds found - nil result",
			mock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"})
				mock.ExpectQuery(regexp.QuoteMeta(dogBreedGetAllQuery)).
					WillReturnRows(rows)
			},
			expectedResult: nil,
			expectedError:  "",
		},
		{
			name: "database error",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(dogBreedGetAllQuery)).
					WillReturnError(errors.New("connection failed"))
			},
			expectedResult: nil,
			expectedError:  "failed to get all dog breeds: connection failed",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("failed to create mock db: %v", err)
			}
			defer db.Close()

			sqlxDB := sqlx.NewDb(db, "sqlmock")
			repo := NewDogBreed(sqlxDB)

			tc.mock(mock)

			result, err := repo.GetAll(context.Background())

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
