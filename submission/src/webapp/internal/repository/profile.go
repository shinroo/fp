package repository

import (
	"github.com/jmoiron/sqlx"
)

type Profile struct {
	db *sqlx.DB
}

func NewProfile(db *sqlx.DB) *Profile {
	return &Profile{db: db}
}
