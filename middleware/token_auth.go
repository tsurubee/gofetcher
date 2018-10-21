package middleware

import (
	"net/http"
	"github.com/tsurubee/gofetcher/config"
)

func TokenAuth(next http.Handler, c *config.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("GOFETCHER-API-TOKEN")

		for _, t := range c.Token {
			if token == t {
				next.ServeHTTP(w, r)
				return
			}
		}
		http.Error(w, "Forbidden", http.StatusForbidden)
	})
}
