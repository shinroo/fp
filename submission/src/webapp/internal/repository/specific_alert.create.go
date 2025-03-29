package repository

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/shinroo/fp/src/webapp/pkg/models"
)

//go:embed queries/specific_alert.create.sql
var specificAlertCreationQuery string

func (r *SpecificAlert) CreateSpecificAlert(ctx context.Context, accountID int, breed string, gender models.Gender, lifeStage models.LifeStage) error {
	res, err := r.db.ExecContext(ctx, specificAlertCreationQuery, accountID, breed, lifeStage, gender)
	if err != nil {
		return fmt.Errorf("failed to insert specific alert: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to insert specific alert: failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("failed to insert specific alert: no rows affected")
	}

	return nil
}
