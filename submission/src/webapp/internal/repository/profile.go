package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/shinroo/fp/src/webapp/pkg/models"
)

type Profile struct {
	db *sqlx.DB
}

func NewProfile(db *sqlx.DB) *Profile {
	return &Profile{db: db}
}

func (r *Profile) CreateProfile(ctx context.Context, accountID int, nearestSPCAID string) error {
	res, err := r.db.ExecContext(ctx, "INSERT INTO profile (account_id, latitude, longitude, nearest_spca) VALUES ($1, 0, 0, $2)", accountID, nearestSPCAID)
	if err != nil {
		return fmt.Errorf("failed to insert profile: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to insert profile: failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("failed to insert profile: no rows affected")
	}

	return nil
}

func (r *Profile) UpdateProfile(ctx context.Context, profile *models.Profile) error {
	res, err := r.db.ExecContext(ctx, `UPDATE profile SET
			latitude = $2, 
			longitude = $3, 
			nearest_spca = $4,
			affectionate_with_family = $5,
			good_with_young_children = $6,
			good_with_other_dogs = $7,
			shedding_level = $8,
			coat_grooming_frequency = $9,
			drooling_level = $10,
			coat_length = $11,
			openness_to_strangers = $12,
			playfulness_level = $13,
			watchdog_protective_nature = $14,
			adaptability_level = $15,
			trainability_level = $16,
			energy_level = $17,
			barking_level = $18,
			mental_stimulation_needs = $19 
		WHERE account_id = $1;`,
		profile.AccountID,
		profile.Latitude,
		profile.Longitude,
		profile.NearestSPCAID,
		profile.AffectionateWithFamily,
		profile.GoodWithYoungChildren,
		profile.GoodWithOtherDogs,
		profile.SheddingLevel,
		profile.CoatGroomingFrequency,
		profile.DroolingLevel,
		profile.CoatLength,
		profile.OpennessToStrangers,
		profile.PlayfulnessLevel,
		profile.WatchdogProtectiveNature,
		profile.AdaptabilityLevel,
		profile.TrainabilityLevel,
		profile.EnergyLevel,
		profile.BarkingLevel,
		profile.MentalStimulationNeeds,
	)
	if err != nil {
		return fmt.Errorf("failed to update profile: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to update profile: failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("failed to update profile: no rows affected")
	}

	return nil
}

func (r *Profile) GetProfileByAccountID(ctx context.Context, account_id int) (*models.Profile, error) {
	var profile models.Profile
	err := r.db.GetContext(ctx, &profile, "SELECT * FROM profile WHERE account_id = $1", account_id)
	if err != nil {
		return nil, fmt.Errorf("failed to get profile by account id: %w", err)
	}

	return &profile, nil
}
