package anywhat

import (
	"net"

	"github.com/dhanusaputra/anywhat-server/api/pb"
	"github.com/dhanusaputra/anywhat-server/pkg/cmd"
	"github.com/dhanusaputra/anywhat-server/pkg/cmd/middleware"
	"github.com/dhanusaputra/anywhat-server/pkg/logger"
	"github.com/dhanusaputra/anywhat-server/pkg/service"
	"github.com/go-playground/validator"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type grpcServer struct {
	anywhat  service.Anywhat
	validate *validator.Validate
}

// GRPCHandler ...
func GRPCHandler(s service.Anywhat, v *validator.Validate, cfg cmd.Config) error {
	lis, err := net.Listen("tcp", ":"+cfg.AnywhatPort)
	if err != nil {
		return err
	}
	opts := []grpc.ServerOption{}
	opts = middleware.AddLogging(logger.Log, opts)
	serv := grpc.NewServer(opts...)
	pb.RegisterAnywhatServer(serv, &grpcServer{s, v})
	// start gRPC server
	logger.Log.Info("starting gRPC server...", zap.String("url", cfg.AnywhatPort))
	return serv.Serve(lis)
}
