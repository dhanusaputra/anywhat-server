package authutil

import (
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/dhanusaputra/anywhat-server/api/pb"
	"github.com/stretchr/testify/assert"
)

func TestSignJWT(t *testing.T) {
	user := &pb.User{
		Id:       "1",
		Username: "username",
		Password: "password",
	}
	tokenString, err := SignJWT(user)
	assert.Nil(t, err)
	assert.NotNil(t, tokenString)
}

func TestValidateJWT(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjpudWxsLCJpZCI6IjEiLCJpc3MiOiJhbnl3aGF0IiwidXNlcm5hbWUiOiJ1c2VybmFtZSJ9.6rNiOIRQ_s5hrDDFw0QdRtmHpz8DC_w22oNRVJzxOiY"
	want := jwt.MapClaims{
		"created_at": nil,
		"id":         "1",
		"iss":        "anywhat",
		"username":   "username",
	}
	_, got, err := ValidateJWT(tokenString)
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}
