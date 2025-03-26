package app

import (
	"log/slog"

	"github.com/shinroo/fp/src/webapp/internal/repository"
)

type Server struct {
	SessionRepo       *repository.Session
	ProfileRepo       *repository.Profile
	DogBreedRepo      *repository.DogBreed
	SpecificAlertRepo *repository.SpecificAlert
	Logger            *slog.Logger
}

func NewServer(
	SessionRepo *repository.Session,
	ProfileRepo *repository.Profile,
	DogBreedRepo *repository.DogBreed,
	SpecificAlertRepo *repository.SpecificAlert,
) *Server {
	return &Server{
		SessionRepo:       SessionRepo,
		ProfileRepo:       ProfileRepo,
		DogBreedRepo:      DogBreedRepo,
		SpecificAlertRepo: SpecificAlertRepo,
		Logger:            slog.Default(),
	}
}
