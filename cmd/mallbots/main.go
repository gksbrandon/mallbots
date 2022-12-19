package main

import (
	"database/sql"
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

	return nil
}
