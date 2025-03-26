package app

import (
	"net/http"

	"github.com/shinroo/fp/src/webapp/internal/keys"
	"github.com/shinroo/fp/src/webapp/pkg/models"
	"golang.org/x/sync/errgroup"
)

func (s *Server) Alerts(w http.ResponseWriter, r *http.Request) {
	session, err := s.SessionRepo.GetSessionByToken(r.Context(), r.URL.Query().Get("sid"))
	if err != nil {
		s.Logger.Error("failed to get session", "error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	errGrp := errgroup.Group{}

	var (
		specificAlerts   []*models.SpecificAlert
		similarityAlerts []*models.SimilarityAlert
		dogBreeds        []*models.DogBreed
	)

	errGrp.Go(func() error {
		var err error
		specificAlerts, err = s.SpecificAlertRepo.GetByAccountID(r.Context(), session.AccountID)
		if err != nil {
			s.Logger.Error("failed to get specific alerts", "error:", err)
			return err
		}
		return nil
	})

	errGrp.Go(func() error {
		var err error
		similarityAlerts, err = s.SimilarityAlertRepo.GetByAccountID(r.Context(), session.AccountID)
		if err != nil {
			s.Logger.Error("failed to get similarity alerts", "error:", err)
			return err
		}
		return nil
	})

	errGrp.Go(func() error {
		var err error
		dogBreeds, err = s.DogBreedRepo.GetAll(r.Context())
		if err != nil {
			s.Logger.Error("failed to get all dog breeds", "error:", err)
			return err
		}
		return nil
	})

	if err := errGrp.Wait(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	renderAppTemplate(w, "alerts", map[string]any{
		"SessionID":        r.Context().Value(keys.SessionID),
		"DogBreeds":        dogBreeds,
		"SpecificAlerts":   specificAlerts,
		"SimilarityAlerts": similarityAlerts,
	})
}
