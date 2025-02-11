package api

import "github.com/shinroo/fp/src/webapp/internal/repository"

type Server struct {
	AccountRepo *repository.Account
	SessionRepo *repository.Session
}

func NewServer(
	AccountRepo *repository.Account,
	SessionRepo *repository.Session,
) *Server {
	return &Server{
		AccountRepo: AccountRepo,
		SessionRepo: SessionRepo,
	}
}
