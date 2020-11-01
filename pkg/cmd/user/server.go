package user

import (
	"context"
	"net"

	"go.uber.org/zap"

	"github.com/dhanusaputra/anywhat-server/api/pb"
	"github.com/dhanusaputra/anywhat-server/pkg/cmd"
	"github.com/dhanusaputra/anywhat-server/pkg/cmd/middleware"
	"github.com/dhanusaputra/anywhat-server/pkg/logger"
	"github.com/dhanusaputra/anywhat-server/pkg/service"
	"google.golang.org/grpc"
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

// GetUser ...
func (s *grpcServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	u, err := s.user.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserResponse{User: u}, nil
}

// ListUser ...
func (s *grpcServer) ListUser(ctx context.Context, _ *emptypb.Empty) (*pb.ListUserResponse, error) {
	us, err := s.user.List(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.ListUserResponse{Users: us}, nil
}

// CreateUser ...
func (s *grpcServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	id, err := s.user.Create(ctx, req.User)
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{Id: id}, nil
}

// UpdateUser ...
func (s *grpcServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	success, err := s.user.Update(ctx, req.User)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserResponse{Updated: success}, nil
}

// DeleteUser ...
func (s *grpcServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	success, err := s.user.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteUserResponse{Deleted: success}, nil
}
