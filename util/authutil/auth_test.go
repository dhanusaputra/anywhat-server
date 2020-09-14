package authutil

import (
	"testing"

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
