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

func TestCreateAccount(t *testing.T) {
	// Create a new mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Wrap the mock database with sqlx
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	// Create an instance of the Account repository with the mocked database
	repo := &Account{db: sqlxDB}

	tests := []struct {
		name          string
		email         string
		password      string
		mockExpect    func()
		expectedError error
	}{
		{
			name:     "successful account creation",
			email:    "test@example.com",
			password: "password123",
			mockExpect: func() {
				// Expect the ExecContext to be called with the correct query and arguments
				mock.ExpectExec("INSERT INTO account \\(email, password\\) VALUES \\(\\$1, \\$2\\)").
					WithArgs("test@example.com", "password123").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name:     "database error",
			email:    "test@example.com",
			password: "password123",
			mockExpect: func() {
				// Simulate a database error
				mock.ExpectExec("INSERT INTO account \\(email, password\\) VALUES \\(\\$1, \\$2\\)").
					WithArgs("test@example.com", "password123").
					WillReturnError(errors.New("database error"))
			},
			expectedError: errors.New("failed to insert account: database error"),
		},
		{
			name:     "no rows affected",
			email:    "test@example.com",
			password: "password123",
			mockExpect: func() {
				// Simulate no rows being affected
				mock.ExpectExec("INSERT INTO account \\(email, password\\) VALUES \\(\\$1, \\$2\\)").
					WithArgs("test@example.com", "password123").
					WillReturnResult(sqlmock.NewResult(0, 0))
			},
			expectedError: errors.New("failed to insert account: no rows affected"),
		},
		{
			name:     "rows affected error",
			email:    "test@example.com",
			password: "password123",
			mockExpect: func() {
				// Simulate an error when getting rows affected
				mock.ExpectExec("INSERT INTO account \\(email, password\\) VALUES \\(\\$1, \\$2\\)").
					WithArgs("test@example.com", "password123").
					WillReturnResult(sqlmock.NewErrorResult(errors.New("rows affected error")))
			},
			expectedError: errors.New("failed to insert account: failed to get rows affected: rows affected error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up the mock expectations
			tt.mockExpect()

			// Call the method under test
			err := repo.CreateAccount(context.Background(), tt.email, tt.password)

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

func TestGetAccountByEmail(t *testing.T) {
	// Create a new mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Wrap the mock database with sqlx
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	// Create an instance of the Account repository with the mocked database
	repo := &Account{db: sqlxDB}

	tests := []struct {
		name           string
		email          string
		mockExpect     func()
		expectedResult *models.Account
		expectedError  error
	}{
		{
			name:  "successful account retrieval",
			email: "test@example.com",
			mockExpect: func() {
				// Mock a successful query with a single row returned
				rows := sqlmock.NewRows([]string{"id", "email", "password"}).
					AddRow(1, "test@example.com", "hashedpassword123")
				mock.ExpectQuery("SELECT \\* FROM account WHERE email = \\$1").
					WithArgs("test@example.com").
					WillReturnRows(rows)
			},
			expectedResult: &models.Account{
				ID:       1,
				Email:    "test@example.com",
				Password: "hashedpassword123",
			},
			expectedError: nil,
		},
		{
			name:  "account not found",
			email: "nonexistent@example.com",
			mockExpect: func() {
				// Mock a query that returns no rows
				mock.ExpectQuery("SELECT \\* FROM account WHERE email = \\$1").
					WithArgs("nonexistent@example.com").
					WillReturnError(sql.ErrNoRows)
			},
			expectedResult: nil,
			expectedError:  fmt.Errorf("failed to get account by email: %w", sql.ErrNoRows),
		},
		{
			name:  "database error",
			email: "test@example.com",
			mockExpect: func() {
				// Mock a database error
				mock.ExpectQuery("SELECT \\* FROM account WHERE email = \\$1").
					WithArgs("test@example.com").
					WillReturnError(errors.New("database error"))
			},
			expectedResult: nil,
			expectedError:  fmt.Errorf("failed to get account by email: %w", errors.New("database error")),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up the mock expectations
			tt.mockExpect()

			// Call the method under test
			result, err := repo.GetAccountByEmail(context.Background(), tt.email)

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
