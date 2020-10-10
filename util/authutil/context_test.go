package authutil

import (
	"context"
	"testing"

	"github.com/dhanusaputra/anywhat-server/api/pb"
	"github.com/stretchr/testify/assert"
)

func TestGetUserContext(t *testing.T) {
	got := GetUserContext(context.WithValue(context.Background(), userCtxKey, &pb.User{}))
	assert.NotNil(t, got)
}

func TestWithUserContext(t *testing.T) {
	got := WithUserContext(context.Background(), &pb.User{})
	assert.NotNil(t, got)
}
