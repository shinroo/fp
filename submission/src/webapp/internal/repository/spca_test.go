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

func TestGetNearest(t *testing.T) {
	// Create a new mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Wrap the mock database with sqlx
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	// Create an instance of the SPCA repository with the mocked database
	repo := NewSPCA(sqlxDB)

	tests := []struct {
		name           string
		latitude       float64
		longitude      float64
		mockExpect     func()
		expectedResult *models.SPCA
		expectedError  error
	}{
		{
			name:      "successful retrieval of nearest SPCA",
			latitude:  37.7749,
			longitude: -122.4194,
			mockExpect: func() {
				// Mock a successful query with a single row returned
				rows := sqlmock.NewRows([]string{"id", "name", "latitude", "longitude", "website", "address"}).
					AddRow("1", "SPCA San Francisco", 37.7749, -122.4194, "https://www.sfspca.org", "123 Main St")
				mock.ExpectQuery(`SELECT id, name, latitude, longitude, website, address FROM spca ORDER BY \(6371 \* acos\(cos\(radians\(\$1\)\) \* cos\(radians\(latitude\)\) \* cos\(radians\(longitude\) - radians\(\$2\)\) \+ sin\(radians\(\$1\)\) \* sin\(radians\(latitude\)\)\)\) LIMIT 1`).
					WithArgs(37.7749, -122.4194).
					WillReturnRows(rows)
			},
			expectedResult: &models.SPCA{
				ID:        "1",
				Name:      "SPCA San Francisco",
				Latitude:  37.7749,
				Longitude: -122.4194,
				Website:   sql.NullString{String: "https://www.sfspca.org", Valid: true},
				Address:   "123 Main St",
			},
			expectedError: nil,
		},
		{
			name:      "no SPCA found",
			latitude:  37.7749,
			longitude: -122.4194,
			mockExpect: func() {
				// Mock a query that returns no rows
				mock.ExpectQuery(`SELECT id, name, latitude, longitude, website, address FROM spca ORDER BY \(6371 \* acos\(cos\(radians\(\$1\)\) \* cos\(radians\(latitude\)\) \* cos\(radians\(longitude\) - radians\(\$2\)\) \+ sin\(radians\(\$1\)\) \* sin\(radians\(latitude\)\)\)\) LIMIT 1`).
					WithArgs(37.7749, -122.4194).
					WillReturnError(sql.ErrNoRows)
			},
			expectedResult: nil,
			expectedError:  fmt.Errorf("failed to get nearest spca: %w", sql.ErrNoRows),
		},
		{
			name:      "database error",
			latitude:  37.7749,
			longitude: -122.4194,
			mockExpect: func() {
				// Mock a database error
				mock.ExpectQuery(`SELECT id, name, latitude, longitude, website, address FROM spca ORDER BY \(6371 \* acos\(cos\(radians\(\$1\)\) \* cos\(radians\(latitude\)\) \* cos\(radians\(longitude\) - radians\(\$2\)\) \+ sin\(radians\(\$1\)\) \* sin\(radians\(latitude\)\)\)\) LIMIT 1`).
					WithArgs(37.7749, -122.4194).
					WillReturnError(errors.New("database error"))
			},
			expectedResult: nil,
			expectedError:  fmt.Errorf("failed to get nearest spca: %w", errors.New("database error")),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up the mock expectations
			tt.mockExpect()

			// Call the method under test
			result, err := repo.GetNearest(context.Background(), tt.latitude, tt.longitude)

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
