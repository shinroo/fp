package repository

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/shinroo/fp/src/webapp/pkg/models"
)

//go:embed queries/specific_alert.get_by_account_id.sql
var specificAlertGetByAccountIDQuery string

func (r *SpecificAlert) GetByAccountID(ctx context.Context, accountID int) ([]*models.SpecificAlert, error) {
	var specificAlerts []*models.SpecificAlert
	err := r.db.SelectContext(ctx, &specificAlerts, specificAlertGetByAccountIDQuery, accountID)
	if err != nil {
		return nil, fmt.Errorf("failed to get specific alerts by account id: %w", err)
	}

	return specificAlerts, nil
}
