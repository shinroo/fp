package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/shinroo/fp/src/webapp/internal/headers"
	"github.com/shinroo/fp/src/webapp/internal/repository"
	"github.com/shinroo/fp/src/webapp/pkg/apimodels"
	"github.com/shinroo/fp/src/webapp/pkg/responses"
)

func authenticateRequest(r *http.Request, sessionRepo *repository.Session) error {
	authHeader := r.Header.Get(headers.AuthKey)
	if authHeader == "" {
		return errors.New("missing auth header")
	}

	userIDHeader := r.Header.Get(headers.UserID)
	if userIDHeader == "" {
		return errors.New("missing user ID header")
	}

	userID, err := strconv.Atoi(userIDHeader)
	if err != nil {
		return fmt.Errorf("failed to convert user ID to int: %w", err)
	}

	session, err := sessionRepo.GetSessionByToken(r.Context(), authHeader)
	if err != nil {
		return fmt.Errorf("failed to get session by token: %w", err)
	}

	if session.AccountID != userID {
		return errors.New("user ID does not match session")
	}

	return nil
}

func WrapWithAuth(next http.Handler, sessionRepo *repository.Session) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := authenticateRequest(r, sessionRepo); err != nil {
			responses.WriteJSON(apimodels.ErrorResponse{
				Message: "Authentication failed",
			}, http.StatusUnauthorized, w)
		}
		next.ServeHTTP(w, r)
	})
}
