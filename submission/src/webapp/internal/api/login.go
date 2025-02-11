package api

import (
	"encoding/json"
	"net/http"

	"github.com/shinroo/fp/src/webapp/pkg/apimodels"
	"github.com/shinroo/fp/src/webapp/pkg/crypto"
	"github.com/shinroo/fp/src/webapp/pkg/responses"
)

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	var req apimodels.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responses.WriteJSON(apimodels.ErrorResponse{
			Message: "failed to unmarshal login request",
		}, http.StatusBadRequest, w)
		return
	}

	if req.Email == "" || req.Password == "" {
		responses.WriteJSON(apimodels.ErrorResponse{
			Message: "email and password are required",
		}, http.StatusBadRequest, w)
		return
	}

	account, err := s.AccountRepo.GetAccountByEmail(r.Context(), req.Email)
	if err != nil {
		responses.WriteJSON(apimodels.ErrorResponse{
			Message: "failed to get account by email",
		}, http.StatusInternalServerError, w)
		return
	}

	err = crypto.ComparePassword(account.Password, req.Password)
	if err != nil {
		responses.WriteJSON(apimodels.ErrorResponse{
			Message: "invalid password",
		}, http.StatusUnauthorized, w)
		return
	}

	token, err := s.SessionRepo.CreateSession(r.Context(), account.ID)
	if err != nil {
		responses.WriteJSON(apimodels.ErrorResponse{
			Message: "failed to create session",
		}, http.StatusInternalServerError, w)
		return
	}

	responses.WriteJSON(apimodels.LoginResponse{
		Token: token,
	}, http.StatusOK, w)
}
