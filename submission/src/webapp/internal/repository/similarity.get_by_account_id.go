package repository

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/shinroo/fp/src/webapp/pkg/models"
)

//go:embed queries/similarity_alert.get_by_account_id.sql
var similarityAlertGetByAccountIDQuery string

func (r *SimilarityAlert) GetByAccountID(ctx context.Context, accountID int) ([]*models.SimilarityAlert, error) {
	var similarityAlerts []*models.SimilarityAlert
	err := r.db.SelectContext(ctx, &similarityAlerts, similarityAlertGetByAccountIDQuery, accountID)
	if err != nil {
		return nil, fmt.Errorf("failed to get similarity alerts by account id: %w", err)
	}

	return similarityAlerts, nil
}
