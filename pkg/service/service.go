package service

import (
	"context"
	"database/sql"

	"github.com/dhanusaputra/anywhat-server/api/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Anywhat ...
type Anywhat interface {
	GetAnything(ctx context.Context, id string) (*pb.Anything, error)
	ListAnything(ctx context.Context) ([]*pb.Anything, error)
	CreateAnything(ctx context.Context, anything *pb.Anything) (string, error)
	UpdateAnything(ctx context.Context, anything *pb.Anything) (bool, error)
	DeleteAnything(ctx context.Context, id string) (bool, error)
}

type anywhatService struct {
	db *sql.DB
}

// NewAnywhatService ...
func NewAnywhatService(db *sql.DB) Anywhat {
	return &anywhatService{db}
}

func (s *anywhatService) GetAnything(ctx context.Context, id string) (*pb.Anything, error) {
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	return nil, nil
}

func (s *anywhatService) ListAnything(ctx context.Context) ([]*pb.Anything, error) {
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	return nil, nil
}

func (s *anywhatService) CreateAnything(ctx context.Context, anything *pb.Anything) (string, error) {
	c, err := s.connect(ctx)
	if err != nil {
		return "", err
	}
	defer c.Close()

	return "", nil
}

func (s *anywhatService) UpdateAnything(ctx context.Context, anything *pb.Anything) (bool, error) {
	c, err := s.connect(ctx)
	if err != nil {
		return false, err
	}
	defer c.Close()

	return false, nil
}

func (s *anywhatService) DeleteAnything(ctx context.Context, id string) (bool, error) {
	c, err := s.connect(ctx)
	if err != nil {
		return false, err
	}
	defer c.Close()

	return false, nil
}

func (s *anywhatService) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := s.db.Conn(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "failed to connect to database: %s", err.Error())
	}
	return c, nil
}
