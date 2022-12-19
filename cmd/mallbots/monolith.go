package main

import (
	"database/sql"
	"mallbots/pkg/config"

	"github.com/rs/zerolog"
)

type monolith struct {
	cfg      config.AppConfig
	logger   zerolog.Logger
	database *sql.DB
}
