package api

import (
	"encoding/json"
	"net/http"

	"github.com/shinroo/fp/src/webapp/pkg/apimodels"
	"github.com/shinroo/fp/src/webapp/pkg/responses"
)

func (s *Server) Signup(w http.ResponseWriter, r *http.Request) {
	var req apimodels.SignupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responses.WriteJSON(apimodels.ErrorResponse{
			Message: "failed to unmarshal signup request",
		}, http.StatusBadRequest, w)
		return
	}

	if req.Email == "" || req.Password == "" {
		responses.WriteJSON(apimodels.ErrorResponse{
			Message: "email and password are required",
		}, http.StatusBadRequest, w)
		return
	}

	err := s.AccountRepo.CreateAccount(r.Context(), req.Email, req.Password)
	if err != nil {
		responses.WriteJSON(apimodels.ErrorResponse{
			Message: "failed to create account",
		}, http.StatusInternalServerError, w)
		return
	}

	responses.WriteJSON(apimodels.SignupResponse{}, http.StatusOK, w)
}
