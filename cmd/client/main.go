package main

import (
	"context"
	"flag"
	"time"

	"github.com/dhanusaputra/anywhat-server/api/pb"
	"github.com/dhanusaputra/anywhat-server/pkg/cmd"
	"github.com/dhanusaputra/anywhat-server/pkg/cmd/anywhat"
	"github.com/dhanusaputra/anywhat-server/pkg/cmd/user"
	"github.com/dhanusaputra/anywhat-server/pkg/logger"
	"github.com/dhanusaputra/anywhat-server/util/envutil"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

var (
	anywhatPort = envutil.GetEnv("ANYWHAT_PORT", "9090")
	userPort    = envutil.GetEnv("USER_PORT", "9091")
)

func main() {
	var cfg cmd.Config
	flag.IntVar(&cfg.LogLevel, "log-level", -1, "Global log level")
	flag.StringVar(&cfg.LogTimeFormat, "log-time-format", "2006-01-02T15:04:05.999999999Z07:00", "Print time format for logger e.g. 006-01-02T15:04:05Z07:00")
	// initialize logger
	if err := logger.Init(cfg.LogLevel, cfg.LogTimeFormat); err != nil {
		panic(err)
	}

	anywhatClient := anywhat.NewClient("localhost:" + anywhatPort)
	userClient := user.NewClient("localhost:" + userPort)
	defer func() {
		anywhatClient.Close()
		userClient.Close()
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	t := time.Now().In(time.UTC)
	tp, err := ptypes.TimestampProto(t)
	if err != nil {
		logger.Log.Fatal(err.Error())
	}

	pfx := t.Format(time.RFC3339Nano)

	req1 := pb.CreateAnythingRequest{
		Anything: &pb.Anything{
			Name:        "name",
			Description: "description (" + pfx + ")",
			CreatedAt:   tp,
			UpdatedAt:   tp,
		},
	}
	res1, err := anywhatClient.Service.CreateAnything(ctx, &req1)
	if err != nil {
		logger.Log.Fatal("CreateAnything failed", zap.Error(err))
	}
	logger.Log.Info("CreateAnything result", zap.Any("res", res1))

	id := res1.Id

	req2 := pb.GetAnythingRequest{
		Id: id,
	}
	res2, err := anywhatClient.Service.GetAnything(ctx, &req2)
	if err != nil {
		logger.Log.Fatal("GetAnything failed", zap.Error(err))
	}
	logger.Log.Info("GetAnything result", zap.Any("res", res2))

	req3 := pb.UpdateAnythingRequest{
		Anything: &pb.Anything{
			Id:          res2.Anything.Id,
			Name:        res2.Anything.Name,
			Description: res2.Anything.Description + " + updated",
			UpdatedAt:   res2.Anything.UpdatedAt,
		},
	}
	res3, err := anywhatClient.Service.UpdateAnything(ctx, &req3)
	if err != nil {
		logger.Log.Fatal("UpdateAnything failed", zap.Error(err))
	}
	logger.Log.Info("UpdateAnything result", zap.Any("res", res3))

	res4, err := anywhatClient.Service.ListAnything(ctx, new(empty.Empty))
	if err != nil {
		logger.Log.Fatal("ListAnything failed", zap.Error(err))
	}
	logger.Log.Info("ListAnything result", zap.Any("res", res4))

	req5 := pb.DeleteAnythingRequest{
		Id: id,
	}
	res5, err := anywhatClient.Service.DeleteAnything(ctx, &req5)
	if err != nil {
		logger.Log.Fatal("DeleteAnything failed", zap.Error(err))
	}
	logger.Log.Info("DeleteAnything result", zap.Any("res", res5))

	req6 := pb.LoginRequest{
		Username: "admin",
		Password: "admin",
	}
	res6, err := userClient.Service.Login(ctx, &req6)
	if err != nil {
		logger.Log.Fatal("Login failed", zap.Error(err))
	}
	logger.Log.Info("Login result", zap.Any("res", res6))

	md := metadata.New(map[string]string{"authorization": res6.Token})
	ctx = metadata.NewOutgoingContext(context.Background(), md)

	res7, err := userClient.Service.Me(ctx, new(empty.Empty))
	if err != nil {
		logger.Log.Fatal("Me failed", zap.Error(err))
	}
	logger.Log.Info("Me result", zap.Any("res", res7))
}
