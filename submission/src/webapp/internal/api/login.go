package api

import (
	"encoding/json"
	"net/http"
)

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Create an empty JSON object
	emptyJSON := struct{}{}

	// Encode the empty JSON object and write it to the response
	json.NewEncoder(w).Encode(emptyJSON)
}
