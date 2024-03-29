package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dhanusaputra/anywhat-server/pkg/cmd"
	"github.com/dhanusaputra/anywhat-server/pkg/cmd/anywhat"
	"github.com/dhanusaputra/anywhat-server/pkg/cmd/user"
	"github.com/dhanusaputra/anywhat-server/pkg/env"
	"github.com/dhanusaputra/anywhat-server/pkg/graph"
	"github.com/dhanusaputra/anywhat-server/pkg/graph/generated"
	"github.com/dhanusaputra/anywhat-server/pkg/graph/middleware"
	"github.com/dhanusaputra/anywhat-server/pkg/logger"
	"github.com/dhanusaputra/anywhat-server/util/envutil"
	"github.com/go-chi/chi"
	"github.com/go-chi/httprate"
	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/zap"
)

var (
	anywhatPort = envutil.GetEnv("ANYWHAT_PORT", "9090")
	userPort    = envutil.GetEnv("USER_PORT", "9091")
	gqlPort     = envutil.GetEnv("GQL_PORT", "3000")
)

func main() {
	var cfg cmd.Config
	flag.IntVar(&cfg.LogLevel, "log-level", -1, "Global log level")
	flag.StringVar(&cfg.LogTimeFormat, "log-time-format", "2006-01-02T15:04:05.999999999Z07:00", "Print time format for logger e.g. 006-01-02T15:04:05Z07:00")
	flag.Parse()

	// init
	if err := logger.Init(cfg.LogLevel, cfg.LogTimeFormat); err != nil {
		panic(err)
	}
	env.Init()

	anywhatClient := anywhat.NewClient("anywhat_server:" + anywhatPort)
	userClient := user.NewClient("user_server:" + userPort)
	defer func() {
		anywhatClient.Close()
		userClient.Close()
	}()

	resolver := graph.NewResolver(anywhatClient.Service, userClient.Service)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	r := chi.NewRouter()

	r.Use(middleware.AddRequestID)
	r.Use(middleware.AddLogger)
	r.Use(middleware.AddAuth)
	r.Use(httprate.Limit(10, 1*time.Minute, httprate.KeyByIP))

	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", srv)

	logger.Log.Info("connect to GraphQL playground", zap.String("host", fmt.Sprintf("http://localhost:%s/", gqlPort)))
	logger.Log.Fatal("listenAndServe failed", zap.Error(http.ListenAndServe(":"+gqlPort, r)))
}
