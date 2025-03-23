package app

import (
	"net/http"

	"github.com/shinroo/fp/src/webapp/internal/keys"
)

func (s *Server) Alerts(w http.ResponseWriter, r *http.Request) {
	renderAppTemplate(w, "alerts", map[string]any{
		"SessionID": r.Context().Value(keys.SessionID),
	})
}
