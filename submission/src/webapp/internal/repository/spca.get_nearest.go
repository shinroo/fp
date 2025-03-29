package repository

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/shinroo/fp/src/webapp/pkg/models"
)

//go:embed queries/spca.get_nearest.sql
var spcaGetNearestQuery string

func (r *SPCA) GetNearest(ctx context.Context, latitude float64, longitude float64) (*models.SPCA, error) {
	var spca models.SPCA
	err := r.db.GetContext(ctx, &spca, spcaGetNearestQuery, latitude, longitude)
	if err != nil {
		return nil, fmt.Errorf("failed to get nearest spca: %w", err)
	}

	return &spca, nil
}
