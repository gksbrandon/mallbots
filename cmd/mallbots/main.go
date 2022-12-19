package main

import (
	"fmt"
	"mallbots/pkg/config"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func run() (err error) {
	var monolith app

	// Parse configuration
	cfg, err := config.InitConfig()
	if err != nil {
		return err
	}
	monolith.cfg = cfg

	return nil
}
