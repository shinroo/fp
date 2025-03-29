package repository

import (
	"github.com/jmoiron/sqlx"
)

type Account struct {
	db *sqlx.DB
}

func NewAccount(db *sqlx.DB) *Account {
	return &Account{db: db}
}
