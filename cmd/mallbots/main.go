package main

import (
	"database/sql"
	"fmt"
	"mallbots/pkg/config"
	"mallbots/pkg/logger"
	"mallbots/pkg/rpc"
	"os"

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

	return nil
}

func initRpc(_ rpc.RpcConfig) *grpc.Server {
	server := grpc.NewServer()

	// Reflection to allow runtime introspection of server protos eg. with gRPCurl
	reflection.Register(server)

	return server
}
