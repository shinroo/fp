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
		s.Logger.Error("email or password were empty", "email:", req.Email)
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

	account, err := s.AccountRepo.GetAccountByEmail(r.Context(), req.Email)
	if err != nil {
		s.Logger.Error("failed to get account by email", "error:", err)
		responses.WriteJSON(apimodels.ErrorResponse{
			Message: "failed to retrieve newly created account",
		}, http.StatusInternalServerError, w)
		return
	}

	spca, err := s.SPCARepo.GetNearest(r.Context(), 0, 0)
	if err != nil {
		s.Logger.Error("failed to find nearest SPCA", "error:", err)
		responses.WriteJSON(apimodels.ErrorResponse{
			Message: "failed to find nearest SPCA",
		}, http.StatusInternalServerError, w)
		return
	}

	err = s.ProfileRepo.CreateProfile(r.Context(), account.ID, spca.ID)
	if err != nil {
		s.Logger.Error("failed to create profile", "error:", err)
		responses.WriteJSON(apimodels.ErrorResponse{
			Message: "failed to create profile",
		}, http.StatusInternalServerError, w)
		return
	}

	s.Logger.Info("signup successful")

	responses.WriteJSON(apimodels.SignupResponse{}, http.StatusOK, w)
}
