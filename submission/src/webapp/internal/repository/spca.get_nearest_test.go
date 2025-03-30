package repository

import (
	"context"
	"database/sql"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/shinroo/fp/src/webapp/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestSPCA_GetNearest(t *testing.T) {
	testSPCA := &models.SPCA{
		ID:        "SPCA123",
		Name:      "Nearest SPCA",
		Latitude:  37.7749,
		Longitude: -122.4194,
		Website:   sql.NullString{String: "https://nearestspca.org", Valid: true},
		Address:   "123 Animal Way, San Francisco, CA",
	}

	testCases := []struct {
		name           string
		latitude       float64
		longitude      float64
		mock           func(mock sqlmock.Sqlmock)
		expectedResult *models.SPCA
		expectedError  string
	}{
		{
			name:      "successful nearest SPCA retrieval",
			latitude:  37.7749,
			longitude: -122.4194,
			mock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name", "latitude", "longitude", "website", "address"}).
					AddRow(testSPCA.ID, testSPCA.Name, testSPCA.Latitude, testSPCA.Longitude, testSPCA.Website, testSPCA.Address)

				mock.ExpectQuery(regexp.QuoteMeta(spcaGetNearestQuery)).
					WithArgs(37.7749, -122.4194).
					WillReturnRows(rows)
			},
			expectedResult: testSPCA,
			expectedError:  "",
		},
		{
			name:      "no SPCA found",
			latitude:  90.0, // North Pole
			longitude: 0.0,
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(spcaGetNearestQuery)).
					WithArgs(90.0, 0.0).
					WillReturnError(sql.ErrNoRows)
			},
			expectedResult: nil,
			expectedError:  "failed to get nearest spca: sql: no rows in result set",
		},
		{
			name:      "database error",
			latitude:  34.0522, // Los Angeles
			longitude: -118.2437,
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(spcaGetNearestQuery)).
					WithArgs(34.0522, -118.2437).
					WillReturnError(errors.New("connection failed"))
			},
			expectedResult: nil,
			expectedError:  "failed to get nearest spca: connection failed",
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
			repo := NewSPCA(sqlxDB)

			tc.mock(mock)

			result, err := repo.GetNearest(context.Background(), tc.latitude, tc.longitude)

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
