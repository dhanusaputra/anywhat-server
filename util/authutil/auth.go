package authutil

import (
	"fmt"
	"time"

	"github.com/dhanusaputra/anywhat-server/api/pb"
	"github.com/dhanusaputra/anywhat-server/pkg/env"
	"github.com/golang-jwt/jwt"
)

const (
	defaultAppName = "anywhat"
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
