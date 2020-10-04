package middleware

import (
	"net/http"

	"github.com/dhanusaputra/anywhat-server/pkg/constant"
	"github.com/dhanusaputra/anywhat-server/util/authutil"
)

// AddAuth ...
func AddAuth() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !constant.AuthEnable {
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
