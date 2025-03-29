package repository

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/shinroo/fp/src/webapp/pkg/models"
)

//go:embed queries/profile.update.sql
var profileUpdateQuery string

func (r *Profile) UpdateProfile(ctx context.Context, profile *models.Profile) error {
	res, err := r.db.ExecContext(
		ctx,
		profileUpdateQuery,
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
