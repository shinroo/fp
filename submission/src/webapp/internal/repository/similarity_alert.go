package repository

import (
	"github.com/jmoiron/sqlx"
)

type SimilarityAlert struct {
	db *sqlx.DB
}

func NewSimilarityAlert(db *sqlx.DB) *SimilarityAlert {
	return &SimilarityAlert{db: db}
}
