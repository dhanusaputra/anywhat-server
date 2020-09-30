package graph

import (
	"github.com/dhanusaputra/anywhat-server/api/pb"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver ...
type Resolver struct {
	anywhatClient pb.AnywhatClient
	userClient    pb.UserServiceClient
}

// NewResolver ...
func NewResolver(anywhatClient pb.AnywhatClient, userClient pb.UserServiceClient) *Resolver {
	return &Resolver{anywhatClient, userClient}
}
