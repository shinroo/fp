package repository

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/shinroo/fp/src/webapp/pkg/models"
)

//go:embed queries/session.get_by_token.sql
var sessionGetByToken string

func (r *Session) GetSessionByToken(ctx context.Context, token string) (*models.Session, error) {
	if token == "" {
		return nil, fmt.Errorf("failed to get session by token: empty token")
	}

	var session models.Session
	err := r.db.GetContext(ctx, &session, sessionGetByToken, token)
	if err != nil {
		return nil, fmt.Errorf("failed to get session by token: %w", err)
	}

	return &session, nil
}
