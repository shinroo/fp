package repository

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/shinroo/fp/src/webapp/pkg/models"
)

//go:embed queries/dog.explore.sql
var dogExploreQuery string

func (r *Dog) Explore(ctx context.Context, similarToVector string) ([]*models.Dog, error) {
	var dogs []*models.Dog

	err := r.db.SelectContext(ctx, &dogs, dogExploreQuery, similarToVector)
	if err != nil {
		return nil, fmt.Errorf("failed to explore dogs: failed to select context: %w", err)
	}

	return dogs, nil
}
