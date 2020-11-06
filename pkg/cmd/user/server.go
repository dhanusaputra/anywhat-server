package user

import (
	"net"

	"go.uber.org/zap"

	"github.com/dhanusaputra/anywhat-server/api/pb"
	"github.com/dhanusaputra/anywhat-server/pkg/cmd"
	"github.com/dhanusaputra/anywhat-server/pkg/cmd/middleware"
	"github.com/dhanusaputra/anywhat-server/pkg/logger"
	"github.com/dhanusaputra/anywhat-server/pkg/service"
	"github.com/go-playground/validator"
	"google.golang.org/grpc"
)

type grpcServer struct {
	user     service.User
	validate *validator.Validate
}

// GRPCHandler ...
func GRPCHandler(s service.User, v *validator.Validate, cfg cmd.Config) error {
	lis, err := net.Listen("tcp", ":"+cfg.UserPort)
	if err != nil {
		return err
	}
	opts := []grpc.ServerOption{}
	opts = middleware.AddLogging(logger.Log, opts)
	serv := grpc.NewServer(opts...)
	pb.RegisterUserServiceServer(serv, &grpcServer{s, v})
	// start gRPC server
	logger.Log.Info("starting gRPC server...", zap.String("url", cfg.UserPort))
	return serv.Serve(lis)
}
