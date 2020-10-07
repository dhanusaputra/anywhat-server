package authutil

import (
	"context"

	"github.com/dhanusaputra/anywhat-server/api/pb"
	"github.com/dhanusaputra/anywhat-server/pkg/logger"
)

type userCtxKeyType string

const userCtxKey userCtxKeyType = "user"

// GetUserContext ...
var GetUserContext = func(ctx context.Context) *pb.User {
	res, ok := ctx.Value(userCtxKey).(*pb.User)
	if !ok {
		logger.Log.Error("convert user failed")
		return nil
	}
	return res
}

// WithUserContext ...
func WithUserContext(ctx context.Context, user *pb.User) context.Context {
	return context.WithValue(ctx, userCtxKey, user)
}
