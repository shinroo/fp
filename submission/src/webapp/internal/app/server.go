package app

import "github.com/shinroo/fp/src/webapp/internal/repository"

type Server struct {
	ProfileRepo *repository.Profile
}

func NewServer(
	ProfileRepo *repository.Profile,
) *Server {
	return &Server{
		ProfileRepo: ProfileRepo,
	}
}
