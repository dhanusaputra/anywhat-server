package service

import (
	"context"
	"time"
)

// Service ...
type Service interface {
	GetAnything(ctx context.Context, id string) (*Anything, error)
	ListAnything(ctx context.Context) ([]*Anything, error)
	CreateAnything(ctx context.Context, anything *Anything) (string, error)
	UpdateAnything(ctx context.Context, anything *Anything) (bool, error)
	DeleteAnything(ctx context.Context, id string) (bool, error)
}

// Anything ...
type Anything struct {
	ID          string
	Name        string
	Description string
	CreatedAt   time.Time
}

type anywhatServer struct{}

// NewService ...
func NewService() Service {
	return &anywhatServer{}
}

func (s anywhatServer) GetAnything(ctx context.Context, id string) (*Anything, error) {
	return nil, nil
}

func (s anywhatServer) ListAnything(ctx context.Context) ([]*Anything, error) {
	return nil, nil
}

func (s anywhatServer) CreateAnything(ctx context.Context, anything *Anything) (string, error) {
	return "", nil
}

func (s anywhatServer) UpdateAnything(ctx context.Context, anything *Anything) (bool, error) {
	return false, nil
}

func (s anywhatServer) DeleteAnything(ctx context.Context, id string) (bool, error) {
	return false, nil
}
