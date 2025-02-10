package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Open(dsn string) (*sqlx.DB, error) {
	return sqlx.Open("postgres", dsn)
}
