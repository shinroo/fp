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

func TestProfile_UpdateProfile(t *testing.T) {
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
		name          string
		profile       *models.Profile
		mock          func(mock sqlmock.Sqlmock)
		expectedError string
	}{
		{
			name:    "successful profile update",
			profile: testProfile,
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(regexp.QuoteMeta(profileUpdateQuery)).
					WithArgs(
						testProfile.AccountID,
						testProfile.Latitude,
						testProfile.Longitude,
						testProfile.NearestSPCAID,
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
					).
					WillReturnResult(sqlmock.NewResult(0, 1)) // For updates, LastInsertId is 0
			},
			expectedError: "",
		},
		{
			name:    "profile not found",
			profile: testProfile,
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(regexp.QuoteMeta(profileUpdateQuery)).
					WithArgs(
						testProfile.AccountID,
						testProfile.Latitude,
						testProfile.Longitude,
						testProfile.NearestSPCAID,
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
					).
					WillReturnResult(sqlmock.NewResult(0, 0))
			},
			expectedError: "failed to update profile: no rows affected",
		},
		{
			name:    "database error",
			profile: testProfile,
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(regexp.QuoteMeta(profileUpdateQuery)).
					WithArgs(
						testProfile.AccountID,
						testProfile.Latitude,
						testProfile.Longitude,
						testProfile.NearestSPCAID,
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
					).
					WillReturnError(errors.New("connection failed"))
			},
			expectedError: "failed to update profile: connection failed",
		},
		{
			name: "invalid SPCA ID format",
			profile: &models.Profile{
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
				NearestSPCAID:            "INVALID_SPCA_ID",
			},
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(regexp.QuoteMeta(profileUpdateQuery)).
					WithArgs(
						123,
						37.7749,
						-122.4194,
						"INVALID_SPCA_ID",
						5, 4, 3, 2, 3, 1, 2, 4, 5, 3, 4, 5, 4, 2, 3).
					WillReturnError(errors.New("invalid spca id format"))
			},
			expectedError: "failed to update profile: invalid spca id format",
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

			err = repo.UpdateProfile(context.Background(), tc.profile)

			if tc.expectedError != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tc.expectedError)
			} else {
				assert.NoError(t, err)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
