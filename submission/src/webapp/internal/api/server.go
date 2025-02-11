package api

import (
	"log/slog"

	"github.com/shinroo/fp/src/webapp/internal/repository"
)

type Server struct {
	Logger      *slog.Logger
	AccountRepo *repository.Account
	SessionRepo *repository.Session
}

func NewServer(
	AccountRepo *repository.Account,
	SessionRepo *repository.Session,
) *Server {
	return &Server{
		Logger:      slog.Default(),
		AccountRepo: AccountRepo,
		SessionRepo: SessionRepo,
	}
}
