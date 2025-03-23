package api

import (
	"log/slog"

	"github.com/shinroo/fp/src/webapp/internal/repository"
)

type Server struct {
	Logger      *slog.Logger
	AccountRepo *repository.Account
	SessionRepo *repository.Session
	ProfileRepo *repository.Profile
	SPCARepo    *repository.SPCA
	DogRepo     *repository.Dog
}

func NewServer(
	AccountRepo *repository.Account,
	SessionRepo *repository.Session,
	ProfileRepo *repository.Profile,
	SPCARepo *repository.SPCA,
	DogRepo *repository.Dog,
) *Server {
	return &Server{
		Logger:      slog.Default(),
		AccountRepo: AccountRepo,
		SessionRepo: SessionRepo,
		ProfileRepo: ProfileRepo,
		SPCARepo:    SPCARepo,
		DogRepo:     DogRepo,
	}
}
