package app

import (
	"net/http"

	"github.com/shinroo/fp/src/webapp/internal/keys"
)

func (s *Server) Search(w http.ResponseWriter, r *http.Request) {
	renderAppTemplate(w, "search", map[string]any{
		"SessionID": r.Context().Value(keys.SessionID),
	})
}
