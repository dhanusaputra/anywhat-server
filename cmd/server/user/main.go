package main

import (
	"database/sql"
	"flag"
	"fmt"

	"github.com/dhanusaputra/anywhat-server/pkg/cmd"
	usercmd "github.com/dhanusaputra/anywhat-server/pkg/cmd/user"
	"github.com/dhanusaputra/anywhat-server/pkg/logger"
	"github.com/dhanusaputra/anywhat-server/pkg/service"
	"github.com/dhanusaputra/anywhat-server/util/envutil"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

var (
	host     = envutil.GetEnv("HOST", "")
	port     = envutil.GetEnvAsInt("PORT", 0)
	user     = envutil.GetEnv("USER", "")
	name     = envutil.GetEnv("NAME", "")
	password = envutil.GetEnv("PASSWORD", "")
)

func main() {
	var cfg cmd.Config
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "9091", "gRPC port to bind")
	flag.IntVar(&cfg.LogLevel, "log-level", -1, "Global log level")
	flag.StringVar(&cfg.LogTimeFormat, "log-time-format", "2006-01-02T15:04:05.999999999Z07:00", "Print time format for logger e.g. 006-01-02T15:04:05Z07:00")
	flag.Parse()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s sslmode=disable",
		host, port, user, password)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// initialize logger
	if err := logger.Init(cfg.LogLevel, cfg.LogTimeFormat); err != nil {
		panic(err)
	}

	s := service.NewUserService(db)
	if err := usercmd.ListenGRPC(s, cfg); err != nil {
		panic(err)
	}
}
