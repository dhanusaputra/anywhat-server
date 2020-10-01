package user

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
	user service.User
}

// ListenGRPC ...
func ListenGRPC(s service.User, cfg cmd.Config) error {
	lis, err := net.Listen("tcp", ":"+cfg.UserPort)
	if err != nil {
		return err
	}
	opts := []grpc.ServerOption{}
	opts = middleware.AddLogging(logger.Log, opts)
	serv := grpc.NewServer(opts...)
	pb.RegisterUserServiceServer(serv, &grpcServer{s})
	reflection.Register(serv)
	// start gRPC server
	logger.Log.Info("starting gRPC server...", zap.String("url", cfg.UserPort))
	return serv.Serve(lis)
}

// Login ...
func (s *grpcServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	token, err := s.user.Login(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	return &pb.LoginResponse{Token: token}, nil
}

// Me ...
func (s *grpcServer) Me(ctx context.Context, _ *emptypb.Empty) (*pb.MeResponse, error) {
	user, err := s.user.Me(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.MeResponse{User: user}, nil
}
