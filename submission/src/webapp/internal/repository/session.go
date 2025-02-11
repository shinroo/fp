package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/shinroo/fp/src/webapp/pkg/models"
)

type Session struct {
	db *sqlx.DB
}

func NewSession(db *sqlx.DB) *Session {
	return &Session{db: db}
}

func (r *Session) CreateSession(ctx context.Context, accountID int) (string, error) {
	token := uuid.NewString()

	res, err := r.db.ExecContext(ctx, "INSERT INTO session (token, account_id) VALUES ($1, $2)", token, accountID)
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

func (r *Session) GetSessionByToken(ctx context.Context, token string) (*models.Session, error) {
	var session models.Session
	err := r.db.GetContext(ctx, &session, "SELECT * FROM session WHERE token = $1", token)
	if err != nil {
		return nil, fmt.Errorf("failed to get session by token: %w", err)
	}

	return &session, nil
}
