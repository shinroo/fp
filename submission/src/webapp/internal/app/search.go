package app

import "net/http"

func (s *Server) Search(w http.ResponseWriter, r *http.Request) {
	renderAppTemplate(w, "search", nil)
}
