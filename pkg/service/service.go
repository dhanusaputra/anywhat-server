package service

import (
	"context"
	"time"
)

// Anywhat ...
type Anywhat interface {
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

type anywhatService struct{}

// NewService ...
func NewService() Anywhat {
	return &anywhatService{}
}

func (s anywhatService) GetAnything(ctx context.Context, id string) (*Anything, error) {
	return nil, nil
}

func (s anywhatService) ListAnything(ctx context.Context) ([]*Anything, error) {
	return nil, nil
}

func (s anywhatService) CreateAnything(ctx context.Context, anything *Anything) (string, error) {
	return "", nil
}

func (s anywhatService) UpdateAnything(ctx context.Context, anything *Anything) (bool, error) {
	return false, nil
}

func (s anywhatService) DeleteAnything(ctx context.Context, id string) (bool, error) {
	return false, nil
}
