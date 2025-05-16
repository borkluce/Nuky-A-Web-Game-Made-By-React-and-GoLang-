package middleware

import (
	"net/http"

	"github.com/kahleryasla/pkg/go/auth/middleware/nethttp"
)

// AuthMiddleware can be used to require authentication for routes
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Use the middleware from the imported package
		nethttp.AuthMiddleware(next).ServeHTTP(w, r)
	})
}
