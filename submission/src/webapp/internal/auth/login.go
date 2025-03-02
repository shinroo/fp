package auth

import "net/http"

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	renderAuthTemplate(w, "login", nil)
}
