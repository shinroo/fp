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

func TestDog_Explore(t *testing.T) {
	testDogs := []*models.Dog{
		{
			Identifier: "dog1",
			Name:       "Buddy",
			Gender:     models.Male,
			LifeStage:  models.Adult,
			ImageUrl:   "http://example.com/dog1.jpg",
			SpcaId:     "SPCA123",
			Breed:      "Labrador",
		},
		{
			Identifier: "dog2",
			Name:       "Bella",
			Gender:     models.Female,
			LifeStage:  models.Puppy,
			ImageUrl:   "http://example.com/dog2.jpg",
			SpcaId:     "SPCA456",
			Breed:      "Beagle",
		},
	}

	testCases := []struct {
		name            string
		similarToVector string
		mock            func(mock sqlmock.Sqlmock)
		expectedResult  []*models.Dog
		expectedError   string
	}{
		{
			name:            "successful exploration with results",
			similarToVector: "vector123",
			mock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{
					"identifier", "name", "gender", "life_stage",
					"image_url", "spca_id", "breed",
				}).
					AddRow(
						testDogs[0].Identifier,
						testDogs[0].Name,
						testDogs[0].Gender,
						testDogs[0].LifeStage,
						testDogs[0].ImageUrl,
						testDogs[0].SpcaId,
						testDogs[0].Breed,
					).
					AddRow(
						testDogs[1].Identifier,
						testDogs[1].Name,
						testDogs[1].Gender,
						testDogs[1].LifeStage,
						testDogs[1].ImageUrl,
						testDogs[1].SpcaId,
						testDogs[1].Breed,
					)

				mock.ExpectQuery(regexp.QuoteMeta(dogExploreQuery)).
					WithArgs("vector123").
					WillReturnRows(rows)
			},
			expectedResult: testDogs,
			expectedError:  "",
		},
		{
			name:            "successful exploration with no results",
			similarToVector: "vector456",
			mock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{
					"identifier", "name", "gender", "life_stage",
					"image_url", "spca_id", "breed",
				})
				mock.ExpectQuery(regexp.QuoteMeta(dogExploreQuery)).
					WithArgs("vector456").
					WillReturnRows(rows)
			},
			expectedResult: nil,
			expectedError:  "",
		},
		{
			name:            "database error",
			similarToVector: "vector789",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(dogExploreQuery)).
					WithArgs("vector789").
					WillReturnError(errors.New("connection failed"))
			},
			expectedResult: nil,
			expectedError:  "failed to explore dogs: failed to select context: connection failed",
		},
		{
			name:            "context canceled",
			similarToVector: "vector999",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(dogExploreQuery)).
					WithArgs("vector999").
					WillReturnError(context.Canceled)
			},
			expectedResult: nil,
			expectedError:  "failed to explore dogs: failed to select context: context canceled",
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
			repo := NewDog(sqlxDB)

			tc.mock(mock)

			result, err := repo.Explore(context.Background(), tc.similarToVector)

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
