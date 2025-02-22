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

func TestCreateProfile(t *testing.T) {
	// Create a new mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Wrap the mock database with sqlx
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	// Create an instance of the Profile repository with the mocked database
	repo := NewProfile(sqlxDB)

	tests := []struct {
		name          string
		accountID     int
		nearestSPCAID string
		mockExpect    func()
		expectedError error
	}{
		{
			name:          "successful profile creation",
			accountID:     123,
			nearestSPCAID: "spca-123",
			mockExpect: func() {
				// Expect the ExecContext to be called with the correct query and arguments
				mock.ExpectExec("INSERT INTO profile \\(account_id, has_children, has_active_lifestyle, has_lots_of_time, latitude, longitude, nearest_spca\\) VALUES \\(\\$1, false, false, false, 0, 0, \\$2\\)").
					WithArgs(123, "spca-123").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name:          "database error",
			accountID:     123,
			nearestSPCAID: "spca-123",
			mockExpect: func() {
				// Simulate a database error
				mock.ExpectExec("INSERT INTO profile \\(account_id, has_children, has_active_lifestyle, has_lots_of_time, latitude, longitude, nearest_spca\\) VALUES \\(\\$1, false, false, false, 0, 0, \\$2\\)").
					WithArgs(123, "spca-123").
					WillReturnError(errors.New("database error"))
			},
			expectedError: fmt.Errorf("failed to insert profile: %w", errors.New("database error")),
		},
		{
			name:          "no rows affected",
			accountID:     123,
			nearestSPCAID: "spca-123",
			mockExpect: func() {
				// Simulate no rows being affected
				mock.ExpectExec("INSERT INTO profile \\(account_id, has_children, has_active_lifestyle, has_lots_of_time, latitude, longitude, nearest_spca\\) VALUES \\(\\$1, false, false, false, 0, 0, \\$2\\)").
					WithArgs(123, "spca-123").
					WillReturnResult(sqlmock.NewResult(0, 0))
			},
			expectedError: fmt.Errorf("failed to insert profile: no rows affected"),
		},
		{
			name:          "rows affected error",
			accountID:     123,
			nearestSPCAID: "spca-123",
			mockExpect: func() {
				// Simulate an error when getting rows affected
				mock.ExpectExec("INSERT INTO profile \\(account_id, has_children, has_active_lifestyle, has_lots_of_time, latitude, longitude, nearest_spca\\) VALUES \\(\\$1, false, false, false, 0, 0, \\$2\\)").
					WithArgs(123, "spca-123").
					WillReturnResult(sqlmock.NewErrorResult(errors.New("rows affected error")))
			},
			expectedError: fmt.Errorf("failed to insert profile: failed to get rows affected: %w", errors.New("rows affected error")),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up the mock expectations
			tt.mockExpect()

			// Call the method under test
			err := repo.CreateProfile(context.Background(), tt.accountID, tt.nearestSPCAID)

			// Assert the expected error
			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}

			// Ensure all expectations were met
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestUpdateProfile(t *testing.T) {
	// Create a new mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Wrap the mock database with sqlx
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	// Create an instance of the Profile repository with the mocked database
	repo := NewProfile(sqlxDB)

	tests := []struct {
		name          string
		profile       models.Profile
		mockExpect    func()
		expectedError error
	}{
		{
			name: "successful profile update",
			profile: models.Profile{
				AccountID:          123,
				HasChildren:        true,
				HasActiveLifestyle: true,
				HasLotsOfTime:      true,
				Latitude:           37.7749,
				Longitude:          -122.4194,
				NearestSPCAID:      "spca-123",
			},
			mockExpect: func() {
				// Expect the ExecContext to be called with the correct query and arguments
				mock.ExpectExec(`UPDATE profile SET
					has_children = \\$2, 
					has_active_lifestyle = \\$3, 
					has_lots_of_time = \\$4, 
					latitude = \\$5, 
					longitude = \\$6, 
					nearest_spca = \\$7 
				WHERE account_id = \\$1`).
					WithArgs(123, true, true, true, 37.7749, -122.4194, "spca-123").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name: "database error",
			profile: models.Profile{
				AccountID:          123,
				HasChildren:        true,
				HasActiveLifestyle: true,
				HasLotsOfTime:      true,
				Latitude:           37.7749,
				Longitude:          -122.4194,
				NearestSPCAID:      "spca-123",
			},
			mockExpect: func() {
				// Simulate a database error
				mock.ExpectExec(`UPDATE profile SET
					has_children = \\$2, 
					has_active_lifestyle = \\$3, 
					has_lots_of_time = \\$4, 
					latitude = \\$5, 
					longitude = \\$6, 
					nearest_spca = \\$7 
				WHERE account_id = \\$1`).
					WithArgs(123, true, true, true, 37.7749, -122.4194, "spca-123").
					WillReturnError(errors.New("database error"))
			},
			expectedError: fmt.Errorf("failed to update profile: %w", errors.New("database error")),
		},
		{
			name: "no rows affected",
			profile: models.Profile{
				AccountID:          123,
				HasChildren:        true,
				HasActiveLifestyle: true,
				HasLotsOfTime:      true,
				Latitude:           37.7749,
				Longitude:          -122.4194,
				NearestSPCAID:      "spca-123",
			},
			mockExpect: func() {
				// Simulate no rows being affected
				mock.ExpectExec(`UPDATE profile SET
					has_children = \\$2, 
					has_active_lifestyle = \\$3, 
					has_lots_of_time = \\$4, 
					latitude = \\$5, 
					longitude = \\$6, 
					nearest_spca = \\$7 
				WHERE account_id = \\$1`).
					WithArgs(123, true, true, true, 37.7749, -122.4194, "spca-123").
					WillReturnResult(sqlmock.NewResult(0, 0))
			},
			expectedError: fmt.Errorf("failed to update profile: no rows affected"),
		},
		{
			name: "rows affected error",
			profile: models.Profile{
				AccountID:          123,
				HasChildren:        true,
				HasActiveLifestyle: true,
				HasLotsOfTime:      true,
				Latitude:           37.7749,
				Longitude:          -122.4194,
				NearestSPCAID:      "spca-123",
			},
			mockExpect: func() {
				// Simulate an error when getting rows affected
				mock.ExpectExec(`UPDATE profile SET
					has_children = \\$2, 
					has_active_lifestyle = \\$3, 
					has_lots_of_time = \\$4, 
					latitude = \\$5, 
					longitude = \\$6, 
					nearest_spca = \\$7 
				WHERE account_id = \\$1`).
					WithArgs(123, true, true, true, 37.7749, -122.4194, "spca-123").
					WillReturnResult(sqlmock.NewErrorResult(errors.New("rows affected error")))
			},
			expectedError: fmt.Errorf("failed to update profile: failed to get rows affected: %w", errors.New("rows affected error")),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up the mock expectations
			tt.mockExpect()

			// Call the method under test
			err := repo.UpdateProfile(context.Background(), tt.profile)

			// Assert the expected error
			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}

			// Ensure all expectations were met
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestGetProfileByAccountID(t *testing.T) {
	// Create a new mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Wrap the mock database with sqlx
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	// Create an instance of the Profile repository with the mocked database
	repo := NewProfile(sqlxDB)

	tests := []struct {
		name           string
		accountID      int
		mockExpect     func()
		expectedResult *models.Profile
		expectedError  error
	}{
		{
			name:      "successful profile retrieval",
			accountID: 123,
			mockExpect: func() {
				// Mock a successful query with a single row returned
				rows := sqlmock.NewRows([]string{"account_id", "has_children", "has_active_lifestyle", "has_lots_of_time", "latitude", "longitude", "nearest_spca"}).
					AddRow(123, true, true, true, 37.7749, -122.4194, "spca-123")
				mock.ExpectQuery("SELECT \\* FROM profile WHERE account_id = \\$1").
					WithArgs(123).
					WillReturnRows(rows)
			},
			expectedResult: &models.Profile{
				AccountID:          123,
				HasChildren:        true,
				HasActiveLifestyle: true,
				HasLotsOfTime:      true,
				Latitude:           37.7749,
				Longitude:          -122.4194,
				NearestSPCAID:      "spca-123",
			},
			expectedError: nil,
		},
		{
			name:      "profile not found",
			accountID: 123,
			mockExpect: func() {
				// Mock a query that returns no rows
				mock.ExpectQuery("SELECT \\* FROM profile WHERE account_id = \\$1").
					WithArgs(123).
					WillReturnError(sql.ErrNoRows)
			},
			expectedResult: nil,
			expectedError:  fmt.Errorf("failed to get profile by account id: %w", sql.ErrNoRows),
		},
		{
			name:      "database error",
			accountID: 123,
			mockExpect: func() {
				// Mock a database error
				mock.ExpectQuery("SELECT \\* FROM profile WHERE account_id = \\$1").
					WithArgs(123).
					WillReturnError(errors.New("database error"))
			},
			expectedResult: nil,
			expectedError:  fmt.Errorf("failed to get profile by account id: %w", errors.New("database error")),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up the mock expectations
			tt.mockExpect()

			// Call the method under test
			result, err := repo.GetProfileByAccountID(context.Background(), tt.accountID)

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
