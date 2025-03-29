package repository

import (
	"context"
	_ "embed"
	"fmt"
)

//go:embed queries/account.create.sql
var accountCreationQuery string

func (r *Account) CreateAccount(ctx context.Context, email string, password string) error {
	res, err := r.db.ExecContext(ctx, accountCreationQuery, email, password)
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
