package repository

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/google/uuid"
)

//go:embed queries/session.create.sql
var sessionCreateQuery string

func (r *Session) CreateSession(ctx context.Context, accountID int) (string, error) {
	token := uuid.NewString()

	res, err := r.db.ExecContext(ctx, sessionCreateQuery, token, accountID)
	if err != nil {
		return "", fmt.Errorf("failed to insert session: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return "", fmt.Errorf("failed to insert session: failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return "", fmt.Errorf("failed to insert session: no rows affected")
	}

	return token, nil
}
