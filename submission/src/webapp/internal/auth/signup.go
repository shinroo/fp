package auth

import "net/http"

func (s *Server) Signup(w http.ResponseWriter, r *http.Request) {
	renderAuthTemplate(w, "signup", nil)
}
