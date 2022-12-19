package main

import (
	"database/sql"
	"mallbots/pkg/config"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

type monolith struct {
	cfg      config.AppConfig
	logger   zerolog.Logger
	database *sql.DB
	rpc      *grpc.Server
}
