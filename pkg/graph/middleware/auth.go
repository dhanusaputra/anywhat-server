package middleware

import (
	"net/http"

	"github.com/dhanusaputra/anywhat-server/pkg/logger"
	"github.com/dhanusaputra/somewhat-server/util/authutil"
	"github.com/dhanusaputra/somewhat-server/util/envutil"
	"go.uber.org/zap"
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
			logger.Log.Info("test", zap.Any("r", r.TransferEncoding))
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
