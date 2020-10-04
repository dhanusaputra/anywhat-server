package middleware

import (
	"net/http"

	"github.com/dhanusaputra/anywhat-server/util/authutil"
	"github.com/dhanusaputra/anywhat-server/util/envutil"
)

const (
	defaultAuthEnable = true
)

// AddAuth ...
func AddAuth() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var (
				authEnable = envutil.GetEnvAsBool("AUTH_ENABLE", defaultAuthEnable)
			)

			if !authEnable {
				next.ServeHTTP(w, r)
				return
			}
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("no authorization found in request"))
				return
			}
			_, _, err := authutil.ValidateJWT(authHeader)
			if err != nil {
				w.WriteHeader(http.StatusForbidden)
				w.Write([]byte(err.Error()))
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
