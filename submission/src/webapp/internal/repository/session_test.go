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

func TestCreateSession(t *testing.T) {
	// Create a new mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Wrap the mock database with sqlx
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	// Create an instance of the Session repository with the mocked database
	repo := &Session{db: sqlxDB}

	tests := []struct {
		name          string
		accountID     int
		mockExpect    func()
		expectedError error
	}{
		{
			name:      "successful session creation",
			accountID: 123,
			mockExpect: func() {
				// Expect the ExecContext to be called with the correct query and arguments
				mock.ExpectExec("INSERT INTO session \\(token, account_id\\) VALUES \\(\\$1, \\$2\\)").
					WithArgs(sqlmock.AnyArg(), 123).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name:      "database error",
			accountID: 123,
			mockExpect: func() {
				// Simulate a database error
				mock.ExpectExec("INSERT INTO session \\(token, account_id\\) VALUES \\(\\$1, \\$2\\)").
					WithArgs(sqlmock.AnyArg(), 123).
					WillReturnError(errors.New("database error"))
			},
			expectedError: fmt.Errorf("failed to insert session: %w", errors.New("database error")),
		},
		{
			name:      "no rows affected",
			accountID: 123,
			mockExpect: func() {
				// Simulate no rows being affected
				mock.ExpectExec("INSERT INTO session \\(token, account_id\\) VALUES \\(\\$1, \\$2\\)").
					WithArgs(sqlmock.AnyArg(), 123).
					WillReturnResult(sqlmock.NewResult(0, 0))
			},
			expectedError: errors.New("failed to insert session: no rows affected"),
		},
		{
			name:      "rows affected error",
			accountID: 123,
			mockExpect: func() {
				// Simulate an error when getting rows affected
				mock.ExpectExec("INSERT INTO session \\(token, account_id\\) VALUES \\(\\$1, \\$2\\)").
					WithArgs(sqlmock.AnyArg(), 123).
					WillReturnResult(sqlmock.NewErrorResult(errors.New("rows affected error")))
			},
			expectedError: fmt.Errorf("failed to insert session: failed to get rows affected: %w", errors.New("rows affected error")),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up the mock expectations
			tt.mockExpect()

			// Call the method under test
			_, err := repo.CreateSession(context.Background(), tt.accountID)

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

func TestGetSessionByToken(t *testing.T) {
	// Create a new mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Wrap the mock database with sqlx
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	// Create an instance of the Session repository with the mocked database
	repo := &Session{db: sqlxDB}

	tests := []struct {
		name           string
		token          string
		mockExpect     func()
		expectedResult *models.Session
		expectedError  error
	}{
		{
			name:  "successful session retrieval",
			token: "token123",
			mockExpect: func() {
				// Mock a successful query with a single row returned
				rows := sqlmock.NewRows([]string{"token", "account_id"}).
					AddRow("token123", "123")
				mock.ExpectQuery("SELECT \\* FROM session WHERE token = \\$1").
					WithArgs("token123").
					WillReturnRows(rows)
			},
			expectedResult: &models.Session{
				Token:     "token123",
				AccountID: 123,
			},
			expectedError: nil,
		},
		{
			name:  "session not found",
			token: "nonexistenttoken",
			mockExpect: func() {
				// Mock a query that returns no rows
				mock.ExpectQuery("SELECT \\* FROM session WHERE token = \\$1").
					WithArgs("nonexistenttoken").
					WillReturnError(sql.ErrNoRows)
			},
			expectedResult: nil,
			expectedError:  fmt.Errorf("failed to get session by token: %w", sql.ErrNoRows),
		},
		{
			name:  "database error",
			token: "token123",
			mockExpect: func() {
				// Mock a database error
				mock.ExpectQuery("SELECT \\* FROM session WHERE token = \\$1").
					WithArgs("token123").
					WillReturnError(errors.New("database error"))
			},
			expectedResult: nil,
			expectedError:  fmt.Errorf("failed to get session by token: %w", errors.New("database error")),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up the mock expectations
			tt.mockExpect()

			// Call the method under test
			result, err := repo.GetSessionByToken(context.Background(), tt.token)

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
