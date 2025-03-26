package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/shinroo/fp/src/webapp/pkg/models"
)

type DogBreed struct {
	db *sqlx.DB
}

func NewDogBreed(db *sqlx.DB) *DogBreed {
	return &DogBreed{db: db}
}

func (r *DogBreed) GetAll(ctx context.Context) ([]*models.DogBreed, error) {
	var dogBreed []*models.DogBreed
	err := r.db.SelectContext(ctx, &dogBreed, `SELECT name FROM dog_breed;`)
	if err != nil {
		return nil, fmt.Errorf("failed to get all dog breeds: %w", err)
	}

	return dogBreed, nil
}
