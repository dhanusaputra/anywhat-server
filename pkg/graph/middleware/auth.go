package middleware

import (
	"context"
	"net/http"

	"github.com/dhanusaputra/anywhat-server/api/pb"
	"github.com/dhanusaputra/anywhat-server/pkg/env"
	"github.com/dhanusaputra/anywhat-server/pkg/logger"
	"github.com/dhanusaputra/anywhat-server/util/authutil"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

// AddAuth ...
func AddAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !env.AuthEnable {
			next.ServeHTTP(w, r)
			return
		}

		header := r.Header.Get("Authorization")
		if header == "" {
			next.ServeHTTP(w, r)
			return
		}

		_, claims, err := authutil.ValidateJWT(header)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(err.Error()))
			return
		}

		ctx := context.WithValue(r.Context(), userCtxKey, &pb.User{
			Id:       claims["id"].(string),
			Username: claims["username"].(string),
		})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// ForContext ...
var ForContext = func(ctx context.Context) *pb.User {
	res, ok := ctx.Value(userCtxKey).(*pb.User)
	if !ok {
		logger.Log.Error("convert user failed")
		return nil
	}
	return res
}
