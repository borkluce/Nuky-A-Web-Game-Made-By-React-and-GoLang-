package nethttp

import (
	// Standart http package
	"net/http"
	"strings"

	"github.com/kahleryasla/pkg/go/auth/token"
	"github.com/kahleryasla/pkg/go/json/util/nethttp"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// 1. Check if the header is set.
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			nethttp.WriteJSON(w, http.StatusUnauthorized, map[string]string{"error": "Authorization header required"})
			return
		}

		// 2. Check if the header starts with "Bearer ".
		// VITAL: (there is a space character after the Bearer)
		if strings.HasPrefix(authHeader, "Bearer ") {
			nethttp.WriteJSON(w, http.StatusUnauthorized, map[string]string{"error": "Invalid authorization type"})
			return
		}
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		// 3. Check if the token is verified.
		_, err := token.VerifyToken(tokenStr)
		if err != nil {
			nethttp.WriteJSON(w, http.StatusUnauthorized, map[string]string{
				"error":   "Invalid token",
				"details": err.Error(),
				"token":   tokenStr,
			})
			return
		}

		// 🟢 SUCESS, go next bro
		next.ServeHTTP(w, r)
	})
}
