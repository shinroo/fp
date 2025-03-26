package app

import (
	"log/slog"

	"github.com/shinroo/fp/src/webapp/internal/repository"
)

type Server struct {
	SessionRepo         *repository.Session
	ProfileRepo         *repository.Profile
	DogBreedRepo        *repository.DogBreed
	SpecificAlertRepo   *repository.SpecificAlert
	SimilarityAlertRepo *repository.SimilarityAlert
	Logger              *slog.Logger
}

func NewServer(
	SessionRepo *repository.Session,
	ProfileRepo *repository.Profile,
	DogBreedRepo *repository.DogBreed,
	SpecificAlertRepo *repository.SpecificAlert,
	SimilarityAlertRepo *repository.SimilarityAlert,
) *Server {
	return &Server{
		SessionRepo:         SessionRepo,
		ProfileRepo:         ProfileRepo,
		DogBreedRepo:        DogBreedRepo,
		SpecificAlertRepo:   SpecificAlertRepo,
		SimilarityAlertRepo: SimilarityAlertRepo,
		Logger:              slog.Default(),
	}
}
