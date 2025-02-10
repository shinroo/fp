package middleware

import "net/http"

func WrapWithAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Implement authentication logic
		next.ServeHTTP(w, r)
	})
}
