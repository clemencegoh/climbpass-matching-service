package middlewares

import (
	"fmt"
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
				return
			}
		}

		// todo: Complete JWT validaity check
		fmt.Println("auth required here")

		// if  {
		// 	http.Error(w, "Unauthorized", http.StatusUnauthorized)
		// }

		next.ServeHTTP(w, r)

		// @After
	})
}
