package main

import (
	"fmt"
	"mallbots/pkg/config"
	"mallbots/pkg/logger"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func run() (err error) {
	var app monolith

	// Parse configuration
	cfg, err := config.InitConfig()
	if err != nil {
		return err
	}
	app.cfg = cfg

	// Initialize logger
	app.logger = logger.New(logger.LogConfig{
		Environment: cfg.Environment,
		LogLevel:    logger.Level(cfg.LogLevel),
	})

	return nil
}
