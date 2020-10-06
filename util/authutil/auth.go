package authutil

import (
	"context"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dhanusaputra/anywhat-server/api/pb"
	"github.com/dhanusaputra/anywhat-server/pkg/env"
	"github.com/dhanusaputra/anywhat-server/pkg/logger"
	"go.uber.org/zap"
)

type ctxKey string

const (
	defaultAppName = "anywhat"

	// CtxKeyUser ...
	CtxKeyUser ctxKey = "auth-user"
)

// SignJWT ...
var SignJWT = func(user *pb.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.Id,
		"username": user.Username,
		"exp":      time.Now().Add(time.Duration(env.JWTExpiredTimeInMinute) * time.Minute).Unix(),
		"iss":      defaultAppName,
	})
	return token.SignedString(env.Key)
}

// ValidateJWT ...
var ValidateJWT = func(tokenString string) (*jwt.Token, jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return env.Key, nil
	})
	return token, claims, err
}

// WithUserContext ...
func WithUserContext(ctx context.Context, user *pb.User) context.Context {
	return context.WithValue(ctx, CtxKeyUser, user)
}

// GetUserContext ...
var GetUserContext = func(ctx context.Context) *pb.User {
	res, ok := ctx.Value(CtxKeyUser).(*pb.User)
	if !ok {
		logger.Log.Error("convert user failed", zap.String("loc", "util.authutil"))
		return nil
	}
	return res
}
