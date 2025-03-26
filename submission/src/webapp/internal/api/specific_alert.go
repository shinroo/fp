package api

import (
	"encoding/json"
	"net/http"

	"github.com/shinroo/fp/src/webapp/pkg/apimodels"
	"github.com/shinroo/fp/src/webapp/pkg/responses"
)

func (s *Server) SpecificAlert(w http.ResponseWriter, r *http.Request) {
	s.Logger.Info("specific alert request received")

	var req apimodels.CreateSpecificAlertRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		s.Logger.Error("failed to unmarshal create specific alert request", "error:", err)
		responses.WriteJSON(apimodels.ErrorResponse{
			Message: "failed to unmarshal create specific alert request",
		}, http.StatusBadRequest, w)
		return
	}

	session, err := s.SessionRepo.GetSessionByToken(r.Context(), r.Header.Get("Authorization"))
	if err != nil {
		s.Logger.Error("failed to get session", "error:", err)
		responses.WriteJSON(apimodels.ErrorResponse{
			Message: "failed to get session",
		}, http.StatusInternalServerError, w)
		return
	}

	err = s.SpecificAlertRepo.CreateSpecificAlert(r.Context(), session.AccountID, req.Breed, req.Gender, req.LifeStage)
	if err != nil {
		s.Logger.Error("failed to create specific alert", "error:", err)
		responses.WriteJSON(apimodels.ErrorResponse{
			Message: "failed to create specific alert",
		}, http.StatusInternalServerError, w)
		return
	}

	s.Logger.Info("creation of specific alert successful")

	responses.WriteJSON(apimodels.CreateSpecificAlertResponse{}, http.StatusOK, w)
}
