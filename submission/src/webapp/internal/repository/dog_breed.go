package repository

import (
	"github.com/jmoiron/sqlx"
)

type DogBreed struct {
	db *sqlx.DB
}

func NewDogBreed(db *sqlx.DB) *DogBreed {
	return &DogBreed{db: db}
}
