package api

import (
	"log/slog"

	"github.com/shinroo/fp/src/webapp/internal/repository"
)

type Server struct {
	Logger              *slog.Logger
	AccountRepo         *repository.Account
	SessionRepo         *repository.Session
	ProfileRepo         *repository.Profile
	SPCARepo            *repository.SPCA
	DogRepo             *repository.Dog
	SpecificAlertRepo   *repository.SpecificAlert
	SimilarityAlertRepo *repository.SimilarityAlert
}

func NewServer(
	AccountRepo *repository.Account,
	SessionRepo *repository.Session,
	ProfileRepo *repository.Profile,
	SPCARepo *repository.SPCA,
	DogRepo *repository.Dog,
	SpecificAlertRepo *repository.SpecificAlert,
	SimilarityAlertRepo *repository.SimilarityAlert,
) *Server {
	return &Server{
		Logger:              slog.Default(),
		AccountRepo:         AccountRepo,
		SessionRepo:         SessionRepo,
		ProfileRepo:         ProfileRepo,
		SPCARepo:            SPCARepo,
		DogRepo:             DogRepo,
		SpecificAlertRepo:   SpecificAlertRepo,
		SimilarityAlertRepo: SimilarityAlertRepo,
	}
}
