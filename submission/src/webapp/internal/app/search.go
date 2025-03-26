package app

import (
	"net/http"

	"github.com/shinroo/fp/src/webapp/internal/keys"
)

func (s *Server) Search(w http.ResponseWriter, r *http.Request) {
	dogBreeds, err := s.DogBreedRepo.GetAll(r.Context())
	if err != nil {
		s.Logger.Error("failed to get all dog breeds", "error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	renderAppTemplate(w, "search", map[string]any{
		"SessionID": r.Context().Value(keys.SessionID),
		"DogBreeds": dogBreeds,
	})
}
