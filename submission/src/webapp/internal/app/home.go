package app

import (
	"net/http"

	"github.com/shinroo/fp/src/webapp/internal/keys"
)

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	renderAppTemplate(w, "home", map[string]any{
		"SessionID": r.Context().Value(keys.SessionID),
	})
}
