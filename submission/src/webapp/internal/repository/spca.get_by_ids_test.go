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

func TestSPCA_GetByIDs(t *testing.T) {
	testSPCAs := []*models.SPCA{
		{
			ID:        "SPCA1",
			Name:      "San Francisco SPCA",
			Latitude:  37.7749,
			Longitude: -122.4194,
			Website:   sql.NullString{String: "https://sfspca.org", Valid: true},
			Address:   "250 Florida St, San Francisco, CA",
		},
		{
			ID:        "SPCA2",
			Name:      "New York SPCA",
			Latitude:  40.7128,
			Longitude: -74.0060,
			Website:   sql.NullString{String: "", Valid: false},
			Address:   "424 E 92nd St, New York, NY",
		},
	}

	testCases := []struct {
		name           string
		spcaIDs        []string
		mock           func(mock sqlmock.Sqlmock, expectedQuery string)
		expectedResult []*models.SPCA
		expectedError  string
	}{
		{
			name:    "successful retrieval with multiple SPCAs",
			spcaIDs: []string{"SPCA1", "SPCA2"},
			mock: func(mock sqlmock.Sqlmock, expectedQuery string) {
				rows := sqlmock.NewRows([]string{"id", "name", "latitude", "longitude", "website", "address"}).
					AddRow(testSPCAs[0].ID, testSPCAs[0].Name, testSPCAs[0].Latitude, testSPCAs[0].Longitude, testSPCAs[0].Website, testSPCAs[0].Address).
					AddRow(testSPCAs[1].ID, testSPCAs[1].Name, testSPCAs[1].Latitude, testSPCAs[1].Longitude, testSPCAs[1].Website, testSPCAs[1].Address)

				mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
					WithArgs("SPCA1", "SPCA2").
					WillReturnRows(rows)
			},
			expectedResult: testSPCAs,
			expectedError:  "",
		},
		{
			name:    "successful retrieval with single SPCA",
			spcaIDs: []string{"SPCA1"},
			mock: func(mock sqlmock.Sqlmock, expectedQuery string) {
				rows := sqlmock.NewRows([]string{"id", "name", "latitude", "longitude", "website", "address"}).
					AddRow(testSPCAs[0].ID, testSPCAs[0].Name, testSPCAs[0].Latitude, testSPCAs[0].Longitude, testSPCAs[0].Website, testSPCAs[0].Address)

				mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
					WithArgs("SPCA1").
					WillReturnRows(rows)
			},
			expectedResult: []*models.SPCA{testSPCAs[0]},
			expectedError:  "",
		},
		{
			name:           "empty IDs array",
			spcaIDs:        []string{},
			mock:           func(_ sqlmock.Sqlmock, _ string) {},
			expectedResult: []*models.SPCA{},
			expectedError:  "",
		},
		{
			name:    "database error",
			spcaIDs: []string{"SPCA1", "SPCA2"},
			mock: func(mock sqlmock.Sqlmock, expectedQuery string) {
				mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
					WithArgs("SPCA1", "SPCA2").
					WillReturnError(errors.New("connection failed"))
			},
			expectedResult: nil,
			expectedError:  "failed to get spca by ids: connection failed",
		},
		{
			name:    "no matching SPCAs",
			spcaIDs: []string{"NON_EXISTENT"},
			mock: func(mock sqlmock.Sqlmock, expectedQuery string) {
				rows := sqlmock.NewRows([]string{"id", "name", "latitude", "longitude", "website", "address"})
				mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
					WithArgs("NON_EXISTENT").
					WillReturnRows(rows)
			},
			expectedResult: nil,
			expectedError:  "",
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

			// Reconstruct the expected query with sqlx.In
			var expectedQuery string
			if len(tc.spcaIDs) > 0 {
				query, _, err := sqlx.In(spcaGetByIDsQuery, tc.spcaIDs)
				if err != nil {
					t.Fatalf("failed to prepare expected query: %v", err)
				}
				expectedQuery = sqlxDB.Rebind(query)
				tc.mock(mock, expectedQuery)
			} else {
				tc.mock(mock, "")
			}

			result, err := repo.GetByIDs(context.Background(), tc.spcaIDs)

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
