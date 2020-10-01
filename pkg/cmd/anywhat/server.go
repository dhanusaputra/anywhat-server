package anywhat

import (
	"context"
	"net"

	"go.uber.org/zap"

	"github.com/dhanusaputra/anywhat-server/api/pb"
	"github.com/dhanusaputra/anywhat-server/pkg/cmd"
	"github.com/dhanusaputra/anywhat-server/pkg/logger"
	"github.com/dhanusaputra/anywhat-server/pkg/middleware"
	"github.com/dhanusaputra/anywhat-server/pkg/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

type grpcServer struct {
	anywhat service.Anywhat
}

// ListenGRPC ...
func ListenGRPC(s service.Anywhat, cfg cmd.Config) error {
	lis, err := net.Listen("tcp", ":"+cfg.AnywhatPort)
	if err != nil {
		return err
	}
	opts := []grpc.ServerOption{}
	opts = middleware.AddLogging(logger.Log, opts)
	serv := grpc.NewServer(opts...)
	pb.RegisterAnywhatServer(serv, &grpcServer{s})
	reflection.Register(serv)
	// start gRPC server
	logger.Log.Info("starting gRPC server...", zap.String("url", cfg.AnywhatPort))
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
