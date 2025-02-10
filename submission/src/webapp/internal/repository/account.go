package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Account struct {
	db *sqlx.DB
}

func (r *Account) CreateAccount(ctx context.Context, email string, password string) error {
	res, err := r.db.ExecContext(ctx, "INSERT INTO accounts (email, password) VALUES ($1, $2)", email, password)
	if err != nil {
		return fmt.Errorf("failed to insert account: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to insert account: failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("failed to insert account: no rows affected")
	}

	return nil
}
