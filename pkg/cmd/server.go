package cmd

import (
	"context"
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
func ListenGRPC(s service.Anywhat, port string) error {
	lis, err := net.Listen("tcp", ":"+port)
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
	a, err := s.anywhat.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetAnythingResponse{Anything: a}, nil
}

// ListAnything ...
func (s *grpcServer) ListAnything(ctx context.Context, _ *emptypb.Empty) (*pb.ListAnythingResponse, error) {
	as, err := s.anywhat.List(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.ListAnythingResponse{Anythings: as}, nil
}

// CreateAnything ...
func (s *grpcServer) CreateAnything(ctx context.Context, req *pb.CreateAnythingRequest) (*pb.CreateAnythingResponse, error) {
	id, err := s.anywhat.Create(ctx, req.Anything)
	if err != nil {
		return nil, err
	}
	return &pb.CreateAnythingResponse{Id: id}, nil
}

// UpdateAnything ...
func (s *grpcServer) UpdateAnything(ctx context.Context, req *pb.UpdateAnythingRequest) (*pb.UpdateAnythingResponse, error) {
	success, err := s.anywhat.Update(ctx, req.Anything)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateAnythingResponse{Updated: success}, nil
}

// DeleteAnything ...
func (s *grpcServer) DeleteAnything(ctx context.Context, req *pb.DeleteAnythingRequest) (*pb.DeleteAnythingResponse, error) {
	success, err := s.anywhat.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteAnythingResponse{Deleted: success}, nil
}
