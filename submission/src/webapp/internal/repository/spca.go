package repository

import (
	"github.com/jmoiron/sqlx"
)

type SPCA struct {
	db *sqlx.DB
}

func NewSPCA(db *sqlx.DB) *SPCA {
	return &SPCA{db: db}
}
