package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/shinroo/fp/src/webapp/pkg/models"
)

type Dog struct {
	db *sqlx.DB
}

func NewDog(db *sqlx.DB) *Dog {
	return &Dog{db: db}
}

type DogSearchFilter interface {
	ToWhereCondition() string
	GetValue() interface{}
}

type DogLifeStageFilter struct {
	LifeStage models.LifeStage
}

func (f DogLifeStageFilter) ToWhereCondition() string {
	return "life_stage = ?"
}

func (f DogLifeStageFilter) GetValue() interface{} {
	return f.LifeStage
}

type DogGenderFilter struct {
	Gender models.Gender
}

func (f DogGenderFilter) ToWhereCondition() string {
	return "gender = ?"
}

func (f DogGenderFilter) GetValue() interface{} {
	return f.Gender
}

type DogBreedFilter struct {
	Breed string
}

func (f DogBreedFilter) ToWhereCondition() string {
	return "breed = ?"
}

func (f DogBreedFilter) GetValue() interface{} {
	return f.Breed
}

func (r *Dog) Search(ctx context.Context, searchKeyword string, filters []DogSearchFilter) ([]*models.Dog, error) {
	query := `SELECT
	dog.identifier AS identifier,
	dog.name AS name,
	dog.gender AS gender,
	dog.life_stage AS life_stage,
	dog.image_url AS image_url,
	spca.id AS spca_id,
	COALESCE(dog.breed, 'Unknown breed') AS breed
	FROM dog
	LEFT JOIN spca ON dog.spca_id = spca.id
	WHERE (dog.name ILIKE '%' || ? || '%'
	OR spca.name ILIKE '%' || ? || '%')`

	args := []interface{}{searchKeyword, searchKeyword}

	if len(filters) > 0 {
		filterConditions := make([]string, len(filters))
		for i, filter := range filters {
			filterConditions[i] = filter.ToWhereCondition()
			args = append(args, filter.GetValue())
		}
		query += " AND " + strings.Join(filterConditions, " AND ")
	}

	var dogs []*models.Dog

	query, args, err := sqlx.In(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare query: %w", err)
	}

	query = r.db.Rebind(query)

	fmt.Println("query:", query)

	err = r.db.SelectContext(ctx, &dogs, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to search for dogs: failed to select context: %w", err)
	}

	return dogs, nil
}

func (r *Dog) Explore(ctx context.Context, similarToVector string) ([]*models.Dog, error) {
	query := `SELECT
	dog.identifier AS identifier,
	dog.name AS name,
	dog.gender AS gender,
	dog.life_stage AS life_stage,
	dog.image_url AS image_url,
	dog.spca_id AS spca_id,
	COALESCE(dog.breed, 'Unknown breed') AS breed
	FROM dog
	ORDER BY dog.embedding <=> $1;`

	var dogs []*models.Dog

	err := r.db.SelectContext(ctx, &dogs, query, similarToVector)
	if err != nil {
		return nil, fmt.Errorf("failed to explore dogs: failed to select context: %w", err)
	}

	return dogs, nil
}
