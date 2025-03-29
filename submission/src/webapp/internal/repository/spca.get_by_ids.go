package repository

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/shinroo/fp/src/webapp/pkg/models"
)

//go:embed queries/spca.get_by_ids.sql
var spcaGetByIDsQuery string

func (r *SPCA) GetByIDs(ctx context.Context, spcaIDs []string) ([]*models.SPCA, error) {
	if len(spcaIDs) == 0 {
		return []*models.SPCA{}, nil
	}

	var spcas []*models.SPCA

	// Use sqlx.In to expand the slice into individual parameters
	query, args, err := sqlx.In(spcaGetByIDsQuery, spcaIDs)
	if err != nil {
		return nil, fmt.Errorf("failed to get spca by ids: failed to prepare query: %w", err)
	}

	// Rebind the query for PostgreSQL (converts ? to $1, $2, etc.)
	query = r.db.Rebind(query)

	err = r.db.SelectContext(ctx, &spcas, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get spca by ids: %w", err)
	}

	return spcas, nil
}
