package repository

import (
	"context"
	_ "embed"
	"fmt"
)

//go:embed queries/profile.create.sql
var profileCreationQuery string

func (r *Profile) CreateProfile(ctx context.Context, accountID int, nearestSPCAID string) error {
	res, err := r.db.ExecContext(ctx, profileCreationQuery, accountID, nearestSPCAID)
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
