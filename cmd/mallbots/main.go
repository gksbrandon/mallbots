package main

import (
	"database/sql"
	"fmt"
	"mallbots/internal/config"
	"mallbots/internal/logger"
	"mallbots/internal/rpc"
	"mallbots/internal/web"
	"os"

	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v4/stdlib"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func run() (err error) {
	// Configure application dependencies
	var app monolith

	// Parse configuration
	app.cfg, err = config.InitConfig()
	if err != nil {
		return err
	}

	// Initialize logger
	app.logger = logger.New(logger.LogConfig{
		Environment: app.cfg.Environment,
		LogLevel:    logger.Level(app.cfg.LogLevel),
	})

	// Initialize database
	app.database, err = sql.Open("pgx", app.cfg.Postgres.Conn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(app.database)

	// Initialize RPC
	app.rpc = initRpc(app.cfg.Rpc)

	// Initialize Mux
	app.mux = initMux(app.cfg.Web)

	return nil
}

func initRpc(_ rpc.RpcConfig) *grpc.Server {
	server := grpc.NewServer()

	// Reflection to allow runtime introspection of server protos eg. with gRPCurl
	reflection.Register(server)

	return server
}

func initMux(_ web.WebConfig) *chi.Mux {
	return chi.NewMux()
}
