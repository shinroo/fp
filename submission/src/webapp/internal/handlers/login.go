package handlers

import "net/http"

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "login", nil)
}
