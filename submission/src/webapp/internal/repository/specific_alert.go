package repository

import (
	"github.com/jmoiron/sqlx"
)

type SpecificAlert struct {
	db *sqlx.DB
}

func NewSpecificAlert(db *sqlx.DB) *SpecificAlert {
	return &SpecificAlert{db: db}
}
