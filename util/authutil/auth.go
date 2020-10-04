package authutil

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dhanusaputra/anywhat-server/api/pb"
	"github.com/dhanusaputra/anywhat-server/util/envutil"
)

var (
	key = []byte(os.Getenv("KEY"))
)

const (
	defaultExpiredTimeInMinute = 15
	defaultAppName             = "anywhat"
)

// SignJWT ...
var SignJWT = func(user *pb.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.Id,
		"username": user.Username,
		"exp":      time.Now().Add(time.Duration(envutil.GetEnvAsInt("JWT_EXPIRED_TIME_IN_MINUTE", defaultExpiredTimeInMinute)) * time.Minute).Unix(),
		"iss":      defaultAppName,
	})
	return token.SignedString(key)
}

// ValidateJWT ...
var ValidateJWT = func(tokenString string) (*jwt.Token, jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})
	return token, claims, err
}

// GetUserContext ...
var GetUserContext = func(ctx context.Context) *pb.User {
	res, _ := ctx.Value(ctx).(*pb.User)
	return res
}
