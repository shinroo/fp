package api

import (
	"encoding/json"
	"net/http"

	"github.com/shinroo/fp/src/webapp/pkg/apimodels"
	"github.com/shinroo/fp/src/webapp/pkg/responses"
)

func (s *Server) SimilarityAlert(w http.ResponseWriter, r *http.Request) {
	s.Logger.Info("similarity alert request received")

	var req apimodels.CreateSimilarityAlertRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		s.Logger.Error("failed to unmarshal create similarity alert request", "error:", err)
		responses.WriteJSON(apimodels.ErrorResponse{
			Message: "failed to unmarshal create similarity alert request",
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

	err = s.SimilarityAlertRepo.CreateSimilarityAlert(r.Context(), session.AccountID, req.SimilarityThreshold)
	if err != nil {
		s.Logger.Error("failed to create similarity alert", "error:", err)
		responses.WriteJSON(apimodels.ErrorResponse{
			Message: "failed to create similarity alert",
		}, http.StatusInternalServerError, w)
		return
	}

	s.Logger.Info("creation of similarity alert successful")

	responses.WriteJSON(apimodels.CreateSimilarityAlertResponse{}, http.StatusOK, w)
}
