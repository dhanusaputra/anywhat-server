package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dhanusaputra/anywhat-server/pkg/cmd"
	"github.com/dhanusaputra/anywhat-server/pkg/cmd/anywhat"
	"github.com/dhanusaputra/anywhat-server/pkg/cmd/user"
	"github.com/dhanusaputra/anywhat-server/pkg/graph"
	"github.com/dhanusaputra/anywhat-server/pkg/graph/generated"
	"github.com/dhanusaputra/anywhat-server/pkg/logger"
	"github.com/dhanusaputra/anywhat-server/util/envutil"
	"go.uber.org/zap"
)

const defaultPort = "3000"

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

	port := envutil.GetEnv("GQL_PORT", defaultPort)

	anywhatClient := anywhat.NewClient("localhost:" + anywhatPort)
	userClient := user.NewClient("localhost:" + userPort)
	defer func() {
		anywhatClient.Close()
		userClient.Close()
	}()

	resolver := graph.NewResolver(anywhatClient.Service, userClient.Service)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	logger.Log.Info("connect to GraphQL playground", zap.String("host", fmt.Sprintf("http://localhost:%s/", port)))
	logger.Log.Fatal("listenAndServe failed", zap.Error(http.ListenAndServe(":"+port, nil)))
}
