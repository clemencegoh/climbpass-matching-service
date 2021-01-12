package middlewares

import (
	"climbpass-matching-service/configs/auth"
	"net/http"
	"strings"
)

// AuthMiddleware does auth for JWT
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// @Before
		excludedPaths := []string{
			"auth",
		}
		for _, b := range excludedPaths {
			if strings.Contains(r.URL.Path, b) {
				next.ServeHTTP(w, r)
			}
		}

		err := auth.TokenValid(r)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)

		// @After
	})
}
