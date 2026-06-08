package middleware

import (
	"net/http"
	"strings"

	"webhook-project/utils"
)

func JWTMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}
		tokenStr := strings.Replace(auth, "Bearer ", "", 1)
		_, err := utils.ValidateToken(tokenStr)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}
