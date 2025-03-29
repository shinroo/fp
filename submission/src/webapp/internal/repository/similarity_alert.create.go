package repository

import (
	"context"
	_ "embed"
	"fmt"
)

//go:embed queries/similarity_alert.create.sql
var similarityAlertCreationQuery string

func (r *SimilarityAlert) CreateSimilarityAlert(ctx context.Context, accountID int, similarityThreshold float32) error {
	res, err := r.db.ExecContext(ctx, similarityAlertCreationQuery, accountID, similarityThreshold)
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
