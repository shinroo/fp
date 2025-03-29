package repository

import (
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
