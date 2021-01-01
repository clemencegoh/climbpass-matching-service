package middlewares

import (
	"net/http"
	"strings"
)

// ContentType does auth for JWT
func ContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// @Before
		excludedPaths := []string{}

		for _, b := range excludedPaths {
			if strings.Contains(r.URL.Path, b) {
				return
			}
		}
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}
