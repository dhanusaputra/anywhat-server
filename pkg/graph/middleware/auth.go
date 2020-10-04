package middleware

import (
	"net/http"

	"github.com/dhanusaputra/somewhat-server/util/authutil"
	"github.com/dhanusaputra/somewhat-server/util/envutil"
)

const (
	defaultAuthEnable = true
)

var (
	authEnable = envutil.GetEnvAsBool("AUTH_ENABLE", defaultAuthEnable)
)

// AddAuth ...
func AddAuth() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
