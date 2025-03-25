package api

import (
	"net/http"

	"github.com/shinroo/fp/src/webapp/pkg/apimodels"
	"github.com/shinroo/fp/src/webapp/pkg/responses"
)

func (s *Server) Explore(w http.ResponseWriter, r *http.Request) {
	s.Logger.Info("explore request received")

	session, err := s.SessionRepo.GetSessionByToken(r.Context(), r.Header.Get("Authorization"))
	if err != nil {
		s.Logger.Error("failed to get session", "error:", err)
		responses.WriteJSON(apimodels.ErrorResponse{
			Message: "failed to get session",
		}, http.StatusInternalServerError, w)
		return
	}

	profile, err := s.ProfileRepo.GetProfileByAccountID(r.Context(), session.AccountID)
	if err != nil {
		s.Logger.Error("failed to get profile", "error:", err)
		responses.WriteJSON(apimodels.ErrorResponse{
			Message: "failed to get profile",
		}, http.StatusInternalServerError, w)
		return
	}

	dogs, err := s.DogRepo.Explore(r.Context(), profile.GetPreferredDogVector())
	if err != nil {
		s.Logger.Error("failed to explore dogs", "error:", err)
		responses.WriteJSON(apimodels.ErrorResponse{
			Message: "failed to explore dogs",
		}, http.StatusInternalServerError, w)
		return
	}

	responses.WriteJSON(apimodels.ExploreResponse{
		Results: dogs,
	}, http.StatusOK, w)
}
