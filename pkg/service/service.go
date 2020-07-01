package service

import (
	"context"

	"github.com/dhanusaputra/anywhat-server/api/pb"
)

// Anywhat ...
type Anywhat interface {
	GetAnything(ctx context.Context, id string) (*pb.Anything, error)
	ListAnything(ctx context.Context) ([]*pb.Anything, error)
	CreateAnything(ctx context.Context, anything *pb.Anything) (string, error)
	UpdateAnything(ctx context.Context, anything *pb.Anything) (bool, error)
	DeleteAnything(ctx context.Context, id string) (bool, error)
}

type anywhatService struct{}

// NewService ...
func NewService() Anywhat {
	return &anywhatService{}
}

func (s anywhatService) GetAnything(ctx context.Context, id string) (*pb.Anything, error) {
	return nil, nil
}

func (s anywhatService) ListAnything(ctx context.Context) ([]*pb.Anything, error) {
	return nil, nil
}

func (s anywhatService) CreateAnything(ctx context.Context, anything *pb.Anything) (string, error) {
	return "", nil
}

func (s anywhatService) UpdateAnything(ctx context.Context, anything *pb.Anything) (bool, error) {
	return false, nil
}

func (s anywhatService) DeleteAnything(ctx context.Context, id string) (bool, error) {
	return false, nil
}
