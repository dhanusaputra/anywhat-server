package main

import (
	"database/sql"
	"flag"
	"fmt"

	"github.com/dhanusaputra/anywhat-server/pkg/cmd"
	"github.com/dhanusaputra/anywhat-server/pkg/service"
	_ "github.com/lib/pq"
)

const (
	host     = "db"
	port     = 5432
	user     = "postgres"
	name     = "postgres"
	password = "postgres"
)

func main() {
	var cfg cmd.Config
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "9090", "gRPC port to bind")
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

	s := service.NewAnywhatService(db)
	if err := cmd.ListenGRPC(s, cfg); err != nil {
		panic(err)
	}
}
