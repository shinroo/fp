package app

import "net/http"

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	renderAppTemplate(w, "home", nil)
}
