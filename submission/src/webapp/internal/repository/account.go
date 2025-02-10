package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/shinroo/fp/src/webapp/pkg/models"
)

type Account struct {
	db *sqlx.DB
}

func NewAccount(db *sqlx.DB) *Account {
	return &Account{db: db}
}

func (r *Account) CreateAccount(ctx context.Context, email string, password string) error {
	res, err := r.db.ExecContext(ctx, "INSERT INTO account (email, password) VALUES ($1, $2)", email, password)
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

func (r *Account) GetAccountByEmail(ctx context.Context, email string) (*models.Account, error) {
	var account models.Account
	err := r.db.GetContext(ctx, &account, "SELECT * FROM account WHERE email = $1", email)
	if err != nil {
		return nil, fmt.Errorf("failed to get account by email: %w", err)
	}

	return &account, nil
}
