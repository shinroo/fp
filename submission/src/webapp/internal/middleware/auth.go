package middleware

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/shinroo/fp/src/webapp/internal/keys"
	"github.com/shinroo/fp/src/webapp/internal/repository"
)

func authenticateRequest(r *http.Request, sessionRepo *repository.Session) error {
	authHeader := r.URL.Query().Get(keys.SessionID)
	if authHeader == "" {
		return errors.New("missing session ID query parameter")
	}

	_, err := sessionRepo.GetSessionByToken(r.Context(), authHeader)
	if err != nil {
		return fmt.Errorf("failed to get session by token: %w", err)
	}

	return nil
}

func WrapWithRedirectAuth(next http.Handler, sessionRepo *repository.Session, logger *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info(
			"authenticating request",
			slog.String("path", r.URL.Path),
			slog.String("sessionID", r.URL.Query().Get(keys.SessionID)),
		)
		if err := authenticateRequest(r, sessionRepo); err != nil {
			logger.Error("failed to authenticate request", slog.Any("error", err))
			http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
			return
		}
		r = r.WithContext(context.WithValue(r.Context(), keys.SessionID, r.URL.Query().Get(keys.SessionID)))
		next.ServeHTTP(w, r)
	})
}
