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
	res, err := r.db.ExecContext(ctx, "INSERT INTO profile (account_id, has_children, has_active_lifestyle, has_lots_of_time, latitude, longitude, nearest_spca) VALUES ($1, false, false, false, 0, 0, $2)", accountID, nearestSPCAID)
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

func (r *Profile) UpdateProfile(ctx context.Context, profile models.Profile) error {
	res, err := r.db.ExecContext(ctx, `UPDATE profile SET
			has_children = $2, 
			has_active_lifestyle = $3, 
			has_lots_of_time = $4, 
			latitude = $5, 
			longitude = $6, 
			nearest_spca = $7 
		WHERE account_id = $1;`,
		profile.AccountID,
		profile.HasChildren,
		profile.HasActiveLifestyle,
		profile.HasLotsOfTime,
		profile.Latitude,
		profile.Longitude,
		profile.NearestSPCAID)
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
