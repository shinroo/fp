package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/shinroo/fp/src/webapp/pkg/models"
)

type SPCA struct {
	db *sqlx.DB
}

func NewSPCA(db *sqlx.DB) *SPCA {
	return &SPCA{db: db}
}

func (r *SPCA) GetNearest(ctx context.Context, latitude float64, longitude float64) (*models.SPCA, error) {
	var spca models.SPCA
	err := r.db.GetContext(ctx, &spca, `SELECT id, name, latitude, longitude, website, address
		FROM spca
		ORDER BY (6371 * acos(cos(radians($1)) * cos(radians(latitude)) * 
		cos(radians(longitude) - radians($2)) + 
		sin(radians($1)) * sin(radians(latitude))))
		LIMIT 1;`, latitude, longitude)
	if err != nil {
		return nil, fmt.Errorf("failed to get nearest spca: %w", err)
	}

	return &spca, nil
}
