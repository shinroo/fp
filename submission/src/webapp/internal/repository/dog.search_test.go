package repository

import (
	"context"
	"errors"
	"regexp"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/shinroo/fp/src/webapp/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestDog_Search(t *testing.T) {
	// Test data
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

	// Test cases
	testCases := []struct {
		name           string
		searchKeyword  string
		filters        []DogSearchFilter
		mock           func(mock sqlmock.Sqlmock, expectedQuery string)
		expectedResult []*models.Dog
		expectedError  string
	}{
		{
			name:          "search with keyword only",
			searchKeyword: "Buddy",
			filters:       nil,
			mock: func(mock sqlmock.Sqlmock, expectedQuery string) {
				rows := sqlmock.NewRows([]string{
					"identifier", "name", "gender", "life_stage",
					"image_url", "spca_id", "breed",
				}).AddRow(
					testDogs[0].Identifier,
					testDogs[0].Name,
					testDogs[0].Gender,
					testDogs[0].LifeStage,
					testDogs[0].ImageUrl,
					testDogs[0].SpcaId,
					testDogs[0].Breed,
				)

				mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
					WithArgs("Buddy", "Buddy").
					WillReturnRows(rows)
			},
			expectedResult: []*models.Dog{testDogs[0]},
			expectedError:  "",
		},
		{
			name:          "search with single filter",
			searchKeyword: "dog",
			filters: []DogSearchFilter{
				DogLifeStageFilter{LifeStage: models.Adult},
			},
			mock: func(mock sqlmock.Sqlmock, expectedQuery string) {
				rows := sqlmock.NewRows([]string{
					"identifier", "name", "gender", "life_stage",
					"image_url", "spca_id", "breed",
				}).AddRow(
					testDogs[0].Identifier,
					testDogs[0].Name,
					testDogs[0].Gender,
					testDogs[0].LifeStage,
					testDogs[0].ImageUrl,
					testDogs[0].SpcaId,
					testDogs[0].Breed,
				)

				mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
					WithArgs("dog", "dog", models.Adult).
					WillReturnRows(rows)
			},
			expectedResult: []*models.Dog{testDogs[0]},
			expectedError:  "",
		},
		{
			name:          "search with multiple filters",
			searchKeyword: "dog",
			filters: []DogSearchFilter{
				DogLifeStageFilter{LifeStage: models.Adult},
				DogGenderFilter{Gender: models.Male},
				DogBreedFilter{Breed: "Labrador"},
			},
			mock: func(mock sqlmock.Sqlmock, expectedQuery string) {
				rows := sqlmock.NewRows([]string{
					"identifier", "name", "gender", "life_stage",
					"image_url", "spca_id", "breed",
				}).AddRow(
					testDogs[0].Identifier,
					testDogs[0].Name,
					testDogs[0].Gender,
					testDogs[0].LifeStage,
					testDogs[0].ImageUrl,
					testDogs[0].SpcaId,
					testDogs[0].Breed,
				)

				mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
					WithArgs("dog", "dog", models.Adult, models.Male, "Labrador").
					WillReturnRows(rows)
			},
			expectedResult: []*models.Dog{testDogs[0]},
			expectedError:  "",
		},
		{
			name:          "no results found",
			searchKeyword: "nonexistent",
			filters:       nil,
			mock: func(mock sqlmock.Sqlmock, expectedQuery string) {
				rows := sqlmock.NewRows([]string{
					"identifier", "name", "gender", "life_stage",
					"image_url", "spca_id", "breed",
				})
				mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
					WithArgs("nonexistent", "nonexistent").
					WillReturnRows(rows)
			},
			expectedResult: nil,
			expectedError:  "",
		},
		{
			name:          "database error",
			searchKeyword: "error",
			filters:       nil,
			mock: func(mock sqlmock.Sqlmock, expectedQuery string) {
				mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
					WithArgs("error", "error").
					WillReturnError(errors.New("connection failed"))
			},
			expectedResult: nil,
			expectedError:  "failed to search for dogs: failed to select context: connection failed",
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
			repo := NewDog(sqlxDB)

			// Build expected query
			query := dogSearchQuery
			if len(tc.filters) > 0 {
				filterConditions := make([]string, len(tc.filters))
				for i, filter := range tc.filters {
					filterConditions[i] = filter.ToWhereCondition()
				}
				query += " AND " + strings.Join(filterConditions, " AND ")
			}

			// Prepare args for sqlx.In
			args := []interface{}{tc.searchKeyword, tc.searchKeyword}
			for _, filter := range tc.filters {
				args = append(args, filter.GetValue())
			}

			// Simulate sqlx.In behavior
			expectedQuery, _, err := sqlx.In(query, args...)
			if err != nil {
				t.Fatalf("failed to prepare expected query: %v", err)
			}
			expectedQuery = sqlxDB.Rebind(expectedQuery)

			// Set up mock expectations
			tc.mock(mock, expectedQuery)

			// Execute the method
			result, err := repo.Search(context.Background(), tc.searchKeyword, tc.filters)

			// Assertions
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
