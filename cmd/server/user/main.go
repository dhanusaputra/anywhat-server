package main

import (
	"database/sql"
	"flag"
	"fmt"

	"github.com/dhanusaputra/anywhat-server/pkg/cmd"
	usercmd "github.com/dhanusaputra/anywhat-server/pkg/cmd/user"
	"github.com/dhanusaputra/anywhat-server/pkg/env"
	"github.com/dhanusaputra/anywhat-server/pkg/logger"
	"github.com/dhanusaputra/anywhat-server/pkg/service"
	"github.com/dhanusaputra/anywhat-server/util/envutil"
	"github.com/go-playground/validator"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

var (
	// db setup
	host     = envutil.GetEnv("DB_USER_HOST", "")
	port     = envutil.GetEnv("PORT", "5432")
	user     = envutil.GetEnv("USER", "")
	name     = envutil.GetEnv("NAME", "")
	password = envutil.GetEnv("PASSWORD", "")

	userPort = envutil.GetEnv("USER_PORT", "9091")
)

func main() {
	var cfg cmd.Config
	flag.StringVar(&cfg.UserPort, "user-port", userPort, "gRPC port to bind")
	flag.IntVar(&cfg.LogLevel, "log-level", -1, "Global log level")
	flag.StringVar(&cfg.LogTimeFormat, "log-time-format", "2006-01-02T15:04:05.999999999Z07:00", "Print time format for logger e.g. 006-01-02T15:04:05Z07:00")
	flag.Parse()

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", host, port, user, password)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// init
	if err := logger.Init(cfg.LogLevel, cfg.LogTimeFormat); err != nil {
		panic(err)
	}
	env.Init()
	v := validator.New()

	s := service.NewUserService(db)
	if err := usercmd.GRPCHandler(s, v, cfg); err != nil {
		panic(err)
	}
}
