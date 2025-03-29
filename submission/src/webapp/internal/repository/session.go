package repository

import (
	"github.com/jmoiron/sqlx"
)

type Session struct {
	db *sqlx.DB
}

func NewSession(db *sqlx.DB) *Session {
	return &Session{db: db}
}
