package middleware

import (
	"context"
	"net/http"
	"strings"

	"services/pkg/auth/token"
)

type contextKey string

// Exported so handlers can access the user info
const UserContextKey = contextKey("user")

// AuthMiddleware validates the Authorization header and attaches user info to the request context
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		// Verify token (adjust based on your VerifyToken implementation)
		claims, err := token.VerifyToken(tokenStr)
		if err != nil {
			http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		// Store claims (or user info) in context using custom key
		ctx := context.WithValue(r.Context(), UserContextKey, claims)

		// Continue with next handler, passing new context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
