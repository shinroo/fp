package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/shinroo/fp/src/webapp/pkg/models"
)

type SimilarityAlert struct {
	db *sqlx.DB
}

func NewSimilarityAlert(db *sqlx.DB) *SimilarityAlert {
	return &SimilarityAlert{db: db}
}

func (r *SimilarityAlert) CreateSimilarityAlert(ctx context.Context, accountID int, similarityThreshold float32) error {
	res, err := r.db.ExecContext(ctx, "INSERT INTO similarity_alert (account_id, similarity_threshold) VALUES ($1, $2)", accountID, similarityThreshold)
	if err != nil {
		return fmt.Errorf("failed to insert similarity alert: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to insert similarity alert: failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("failed to insert similarity alert: no rows affected")
	}

	return nil
}

func (r *SimilarityAlert) GetByAccountID(ctx context.Context, accountID int) ([]*models.SimilarityAlert, error) {
	var similarityAlerts []*models.SimilarityAlert
	err := r.db.SelectContext(ctx, &similarityAlerts, "SELECT * FROM similarity_alert WHERE account_id = $1", accountID)
	if err != nil {
		return nil, fmt.Errorf("failed to get similarity alerts by account id: %w", err)
	}

	return similarityAlerts, nil
}
