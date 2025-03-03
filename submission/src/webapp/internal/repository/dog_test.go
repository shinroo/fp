package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/shinroo/fp/src/webapp/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestDogSearch(t *testing.T) {
	// Create a new mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Wrap the mock database with sqlx
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	// Create an instance of the Dog repository with the mocked database
	repo := NewDog(sqlxDB)

	tests := []struct {
		name           string
		searchKeyword  string
		filters        []DogSearchFilter
		mockExpect     func()
		expectedResult []*models.Dog
		expectedError  error
	}{
		{
			name:          "successful search without filters",
			searchKeyword: "Buddy",
			filters:       nil,
			mockExpect: func() {
				// Mock a successful query with a single row returned
				rows := sqlmock.NewRows([]string{"identifier", "name", "gender", "life_stage", "image_url", "spca_id"}).
					AddRow("dog-123", "Buddy", models.Male, models.Adult, "https://example.com/buddy.jpg", "spca-1")
				mock.ExpectQuery("SELECT \\* FROM dog LEFT JOIN spca ON dog.spca_id = spca.id WHERE dog.name ILIKE '%' || \\? || '%' OR spca.name ILIKE '%' || \\? || '%'").
					WithArgs("Buddy", "Buddy").
					WillReturnRows(rows)
			},
			expectedResult: []*models.Dog{
				{
					Identifier: "dog-123",
					Name:       "Buddy",
					Gender:     models.Male,
					LifeStage:  models.Adult,
					ImageUrl:   "https://example.com/buddy.jpg",
					SpcaId:     "spca-1",
				},
			},
			expectedError: nil,
		},
		{
			name:          "successful search with life stage filter",
			searchKeyword: "Buddy",
			filters: []DogSearchFilter{
				DogLifeStageFilter{LifeStage: models.Adult},
			},
			mockExpect: func() {
				// Mock a successful query with a single row returned
				rows := sqlmock.NewRows([]string{"identifier", "name", "gender", "life_stage", "image_url", "spca_id"}).
					AddRow("dog-123", "Buddy", models.Male, models.Adult, "https://example.com/buddy.jpg", "spca-1")
				mock.ExpectQuery("SELECT \\* FROM dog LEFT JOIN spca ON dog.spca_id = spca.id WHERE dog.name ILIKE '%' || \\? || '%' OR spca.name ILIKE '%' || \\? || '%' AND life_stage = \\?").
					WithArgs("Buddy", "Buddy", string(models.Adult)).
					WillReturnRows(rows)
			},
			expectedResult: []*models.Dog{
				{
					Identifier: "dog-123",
					Name:       "Buddy",
					Gender:     models.Male,
					LifeStage:  models.Adult,
					ImageUrl:   "https://example.com/buddy.jpg",
					SpcaId:     "spca-1",
				},
			},
			expectedError: nil,
		},
		{
			name:          "successful search with gender filter",
			searchKeyword: "Buddy",
			filters: []DogSearchFilter{
				DogGenderFilter{Gender: models.Male},
			},
			mockExpect: func() {
				// Mock a successful query with a single row returned
				rows := sqlmock.NewRows([]string{"identifier", "name", "gender", "life_stage", "image_url", "spca_id"}).
					AddRow("dog-123", "Buddy", models.Male, models.Adult, "https://example.com/buddy.jpg", "spca-1")
				mock.ExpectQuery("SELECT \\* FROM dog LEFT JOIN spca ON dog.spca_id = spca.id WHERE dog.name ILIKE '%' || \\? || '%' OR spca.name ILIKE '%' || \\? || '%' AND gender = \\?").
					WithArgs("Buddy", "Buddy", string(models.Male)).
					WillReturnRows(rows)
			},
			expectedResult: []*models.Dog{
				{
					Identifier: "dog-123",
					Name:       "Buddy",
					Gender:     models.Male,
					LifeStage:  models.Adult,
					ImageUrl:   "https://example.com/buddy.jpg",
					SpcaId:     "spca-1",
				},
			},
			expectedError: nil,
		},
		{
			name:          "successful search with multiple filters",
			searchKeyword: "Buddy",
			filters: []DogSearchFilter{
				DogLifeStageFilter{LifeStage: models.Adult},
				DogGenderFilter{Gender: models.Male},
			},
			mockExpect: func() {
				// Mock a successful query with a single row returned
				rows := sqlmock.NewRows([]string{"identifier", "name", "gender", "life_stage", "image_url", "spca_id"}).
					AddRow("dog-123", "Buddy", models.Male, models.Adult, "https://example.com/buddy.jpg", "spca-1")
				mock.ExpectQuery("SELECT \\* FROM dog LEFT JOIN spca ON dog.spca_id = spca.id WHERE dog.name ILIKE '%' || \\? || '%' OR spca.name ILIKE '%' || \\? || '%' AND life_stage = \\? AND gender = \\?").
					WithArgs("Buddy", "Buddy", string(models.Adult), string(models.Male)).
					WillReturnRows(rows)
			},
			expectedResult: []*models.Dog{
				{
					Identifier: "dog-123",
					Name:       "Buddy",
					Gender:     models.Male,
					LifeStage:  models.Adult,
					ImageUrl:   "https://example.com/buddy.jpg",
					SpcaId:     "spca-1",
				},
			},
			expectedError: nil,
		},
		{
			name:          "database error",
			searchKeyword: "Buddy",
			filters:       nil,
			mockExpect: func() {
				// Mock a database error
				mock.ExpectQuery("SELECT \\* FROM dog LEFT JOIN spca ON dog.spca_id = spca.id WHERE dog.name ILIKE '%' || \\? || '%' OR spca.name ILIKE '%' || \\? || '%'").
					WithArgs("Buddy", "Buddy").
					WillReturnError(errors.New("database error"))
			},
			expectedResult: nil,
			expectedError:  fmt.Errorf("failed to search for dogs: failed to select context: %w", errors.New("database error")),
		},
		{
			name:          "no rows found",
			searchKeyword: "Unknown",
			filters:       nil,
			mockExpect: func() {
				// Mock a query that returns no rows
				mock.ExpectQuery("SELECT \\* FROM dog LEFT JOIN spca ON dog.spca_id = spca.id WHERE dog.name ILIKE '%' || \\? || '%' OR spca.name ILIKE '%' || \\? || '%'").
					WithArgs("Unknown", "Unknown").
					WillReturnError(sql.ErrNoRows)
			},
			expectedResult: nil,
			expectedError:  fmt.Errorf("failed to search for dogs: failed to select context: %w", sql.ErrNoRows),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up the mock expectations
			tt.mockExpect()

			// Call the method under test
			result, err := repo.Search(context.Background(), tt.searchKeyword, tt.filters)

			// Assert the expected result and error
			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.expectedResult, result)

			// Ensure all expectations were met
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
