package repository

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/shinroo/fp/src/webapp/pkg/models"
)

//go:embed queries/dog_breed.get_all.sql
var dogBreedGetAllQuery string

func (r *DogBreed) GetAll(ctx context.Context) ([]*models.DogBreed, error) {
	var dogBreed []*models.DogBreed
	err := r.db.SelectContext(ctx, &dogBreed, dogBreedGetAllQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to get all dog breeds: %w", err)
	}

	return dogBreed, nil
}
