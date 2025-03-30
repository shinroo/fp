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

func TestProfile_GetProfileByAccountID(t *testing.T) {
	testProfile := &models.Profile{
		AccountID:                123,
		AffectionateWithFamily:   5,
		GoodWithYoungChildren:    4,
		GoodWithOtherDogs:        3,
		SheddingLevel:            2,
		CoatGroomingFrequency:    3,
		DroolingLevel:            1,
		CoatLength:               2,
		OpennessToStrangers:      4,
		PlayfulnessLevel:         5,
		WatchdogProtectiveNature: 3,
		AdaptabilityLevel:        4,
		TrainabilityLevel:        5,
		EnergyLevel:              4,
		BarkingLevel:             2,
		MentalStimulationNeeds:   3,
		Latitude:                 37.7749,
		Longitude:                -122.4194,
		NearestSPCAID:            "SPCA123",
	}

	testCases := []struct {
		name           string
		accountID      int
		mock           func(mock sqlmock.Sqlmock)
		expectedResult *models.Profile
		expectedError  string
	}{
		{
			name:      "successful profile retrieval",
			accountID: 123,
			mock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{
					"account_id", "affectionate_with_family", "good_with_young_children",
					"good_with_other_dogs", "shedding_level", "coat_grooming_frequency",
					"drooling_level", "coat_length", "openness_to_strangers",
					"playfulness_level", "watchdog_protective_nature", "adaptability_level",
					"trainability_level", "energy_level", "barking_level",
					"mental_stimulation_needs", "latitude", "longitude", "nearest_spca",
				}).
					AddRow(
						testProfile.AccountID,
						testProfile.AffectionateWithFamily,
						testProfile.GoodWithYoungChildren,
						testProfile.GoodWithOtherDogs,
						testProfile.SheddingLevel,
						testProfile.CoatGroomingFrequency,
						testProfile.DroolingLevel,
						testProfile.CoatLength,
						testProfile.OpennessToStrangers,
						testProfile.PlayfulnessLevel,
						testProfile.WatchdogProtectiveNature,
						testProfile.AdaptabilityLevel,
						testProfile.TrainabilityLevel,
						testProfile.EnergyLevel,
						testProfile.BarkingLevel,
						testProfile.MentalStimulationNeeds,
						testProfile.Latitude,
						testProfile.Longitude,
						testProfile.NearestSPCAID,
					)

				mock.ExpectQuery(regexp.QuoteMeta(profileGetByAccountIDQuery)).
					WithArgs(123).
					WillReturnRows(rows)
			},
			expectedResult: testProfile,
			expectedError:  "",
		},
		{
			name:      "profile not found",
			accountID: 999,
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(profileGetByAccountIDQuery)).
					WithArgs(999).
					WillReturnError(sql.ErrNoRows)
			},
			expectedResult: nil,
			expectedError:  "failed to get profile by account id: sql: no rows in result set",
		},
		{
			name:      "database error",
			accountID: 456,
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(profileGetByAccountIDQuery)).
					WithArgs(456).
					WillReturnError(errors.New("connection failed"))
			},
			expectedResult: nil,
			expectedError:  "failed to get profile by account id: connection failed",
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
			repo := NewProfile(sqlxDB)

			tc.mock(mock)

			result, err := repo.GetProfileByAccountID(context.Background(), tc.accountID)

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
