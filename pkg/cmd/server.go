package cmd

import (
	"context"
	"fmt"
	"net"

	"github.com/dhanusaputra/anywhat-server/api/pb"
	"github.com/dhanusaputra/anywhat-server/pkg/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

type grpcServer struct {
	anywhat service.Anywhat
}

// ListenGRPC ...
func ListenGRPC(s service.Anywhat, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	serv := grpc.NewServer()
	pb.RegisterAnywhatServer(serv, &grpcServer{s})
	reflection.Register(serv)
	return serv.Serve(lis)
}

// GetAnything ...
func (s *grpcServer) GetAnything(ctx context.Context, req *pb.GetAnythingRequest) (*pb.GetAnythingResponse, error) {
	_, err := s.anywhat.GetAnything(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetAnythingResponse{}, nil
}

// ListAnything ...
func (s *grpcServer) ListAnything(ctx context.Context, _ *emptypb.Empty) (*pb.ListAnythingResponse, error) {
	_, err := s.anywhat.ListAnything(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.ListAnythingResponse{}, nil
}

// CreateAnything ...
func (s *grpcServer) CreateAnything(ctx context.Context, req *pb.CreateAnythingRequest) (*pb.CreateAnythingResponse, error) {
	_, err := s.anywhat.CreateAnything(ctx, req.Anything)
	if err != nil {
		return nil, err
	}
	return &pb.CreateAnythingResponse{}, nil
}

// UpdateAnything ...
func (s *grpcServer) UpdateAnything(ctx context.Context, req *pb.UpdateAnythingRequest) (*pb.UpdateAnythingResponse, error) {
	_, err := s.anywhat.UpdateAnything(ctx, req.Anything)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateAnythingResponse{}, nil
}

// DeleteAnything ...
func (s *grpcServer) DeleteAnything(ctx context.Context, req *pb.DeleteAnythingRequest) (*pb.DeleteAnythingResponse, error) {
	_, err := s.anywhat.DeleteAnything(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteAnythingResponse{}, nil
}
