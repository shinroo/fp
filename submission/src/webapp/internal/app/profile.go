package app

import "net/http"

func (s *Server) Profile(w http.ResponseWriter, r *http.Request) {
	renderAppTemplate(w, "profile", nil)
}
