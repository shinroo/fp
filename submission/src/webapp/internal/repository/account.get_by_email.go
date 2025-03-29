package repository

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/shinroo/fp/src/webapp/pkg/models"
)

//go:embed queries/account.get_by_email.sql
var accountGetByEmailQuery string

func (r *Account) GetAccountByEmail(ctx context.Context, email string) (*models.Account, error) {
	var account models.Account
	err := r.db.GetContext(ctx, &account, accountGetByEmailQuery, email)
	if err != nil {
		return nil, fmt.Errorf("failed to get account by email: %w", err)
	}

	return &account, nil
}
