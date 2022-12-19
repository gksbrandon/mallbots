package main

import (
	"mallbots/pkg/config"

	"github.com/rs/zerolog"
)

type monolith struct {
	cfg    config.AppConfig
	logger zerolog.Logger
}
