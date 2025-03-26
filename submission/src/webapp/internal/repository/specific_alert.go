package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/shinroo/fp/src/webapp/pkg/models"
)

type SpecificAlert struct {
	db *sqlx.DB
}

func NewSpecificAlert(db *sqlx.DB) *SpecificAlert {
	return &SpecificAlert{db: db}
}

func (r *SpecificAlert) CreateSpecificAlert(ctx context.Context, accountID int, breed string, gender models.Gender, lifeStage models.LifeStage) error {
	res, err := r.db.ExecContext(ctx, "INSERT INTO specific_alert (account_id, breed, life_stage, gender) VALUES ($1, $2, $3, $4)", accountID, breed, lifeStage, gender)
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

func (r *SpecificAlert) GetByAccountID(ctx context.Context, accountID int) ([]*models.SpecificAlert, error) {
	var specificAlerts []*models.SpecificAlert
	err := r.db.SelectContext(ctx, &specificAlerts, "SELECT * FROM specific_alert WHERE account_id = $1", accountID)
	if err != nil {
		return nil, fmt.Errorf("failed to get specific alerts by account id: %w", err)
	}

	return specificAlerts, nil
}
