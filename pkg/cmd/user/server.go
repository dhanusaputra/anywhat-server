package user

import (
	"context"
	"fmt"
	"net"

	"go.uber.org/zap"

	"github.com/dhanusaputra/anywhat-server/api/pb"
	"github.com/dhanusaputra/anywhat-server/pkg/logger"
	"github.com/dhanusaputra/anywhat-server/pkg/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

type grpcServer struct {
	user service.User
}

// Config is configuration for Server
type Config struct {
	// gRPC server start parameters section
	// gRPC is TCP port to listen by gRPC server
	GRPCPort string

	// Log parameters section
	// LogLevel is global log level: Debug(-1), Info(0), Warn(1), Error(2), DPanic(3), Panic(4), Fatal(5)
	LogLevel int
	// LogTimeFormat is print time format for logger e.g. 2006-01-02T15:04:05Z07:00
	LogTimeFormat string
}

// ListenGRPC ...
func ListenGRPC(s service.User, cfg Config) error {
	lis, err := net.Listen("tcp", ":"+cfg.GRPCPort)
	if err != nil {
		return err
	}
	// initialize logger
	if err := logger.Init(cfg.LogLevel, cfg.LogTimeFormat); err != nil {
		return fmt.Errorf("failed to initialize logger: %v", err)
	}
	serv := grpc.NewServer()
	pb.RegisterUserServiceServer(serv, &grpcServer{s})
	reflection.Register(serv)
	// start gRPC server
	logger.Log.Info("starting gRPC server...", zap.String("port", cfg.GRPCPort))
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
