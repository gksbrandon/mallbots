package main

import (
	"database/sql"
	"mallbots/internal/config"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

type monolith struct {
	cfg      config.AppConfig
	logger   zerolog.Logger
	database *sql.DB
	rpc      *grpc.Server
	mux      *chi.Mux
}
