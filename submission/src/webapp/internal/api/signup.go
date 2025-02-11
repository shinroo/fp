package api

import (
	"encoding/json"
	"net/http"

	"github.com/shinroo/fp/src/webapp/pkg/apimodels"
	"github.com/shinroo/fp/src/webapp/pkg/crypto"
	"github.com/shinroo/fp/src/webapp/pkg/responses"
)

func (s *Server) Signup(w http.ResponseWriter, r *http.Request) {
	s.Logger.Info("signup request received")

	var req apimodels.SignupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		s.Logger.Error("failed to unmarshal signup request", "error:", err)
		responses.WriteJSON(apimodels.ErrorResponse{
			Message: "failed to unmarshal signup request",
		}, http.StatusBadRequest, w)
		return
	}

	if req.Email == "" || req.Password == "" {
		s.Logger.Error("email or password were empty", "email:", req.Email, "password:", req.Password)
		responses.WriteJSON(apimodels.ErrorResponse{
			Message: "email and password are required",
		}, http.StatusBadRequest, w)
		return
	}

	passwordHash, err := crypto.HashPassword(req.Password)
	if err != nil {
		s.Logger.Error("failed to hash password", "error:", err)
		responses.WriteJSON(apimodels.ErrorResponse{
			Message: "signup failed due to an unexpected error",
		}, http.StatusInternalServerError, w)
		return
	}

	err = s.AccountRepo.CreateAccount(r.Context(), req.Email, passwordHash)
	if err != nil {
		s.Logger.Error("failed to create a new account", "error:", err)
		responses.WriteJSON(apimodels.ErrorResponse{
			Message: "failed to create account",
		}, http.StatusInternalServerError, w)
		return
	}

	s.Logger.Info("signup successful")

	responses.WriteJSON(apimodels.SignupResponse{}, http.StatusOK, w)
}
