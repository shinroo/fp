package repository

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/shinroo/fp/src/webapp/pkg/models"
)

//go:embed queries/profile.get_by_account_id.sql
var profileGetByAccountIDQuery string

func (r *Profile) GetProfileByAccountID(ctx context.Context, account_id int) (*models.Profile, error) {
	var profile models.Profile
	err := r.db.GetContext(ctx, &profile, profileGetByAccountIDQuery, account_id)
	if err != nil {
		return nil, fmt.Errorf("failed to get profile by account id: %w", err)
	}

	return &profile, nil
}
