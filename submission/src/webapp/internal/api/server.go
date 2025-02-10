package api

import "github.com/shinroo/fp/src/webapp/internal/repository"

type Server struct {
	AccountRepo *repository.Account
}

func NewServer(
	AccountRepo *repository.Account,
) *Server {
	return &Server{
		AccountRepo: AccountRepo,
	}
}
