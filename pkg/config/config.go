package config

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/stackus/dotenv"
)

type (
	PGConfig struct {
		Conn string `required:"true"`
	}

	AppConfig struct {
		Environment string
		LogLevel    string `envconfig:"LOG_LEVEL" default:"DEBUG"`
	}
)

func InitConfig() (cfg AppConfig, err error) {
	if err = dotenv.Load(dotenv.EnvironmentFiles(os.Getenv("ENVIRONMENT"))); err != nil {
		return
	}

	err = envconfig.Process("", &cfg)

	return
}
