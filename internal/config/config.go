package config

import (
	"mallbots/internal/rpc"
	"mallbots/internal/web"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/stackus/dotenv"
)

type (
	PostgresConfig struct {
		Conn string `required:"true"`
	}

	AppConfig struct {
		Environment string
		LogLevel    string `envconfig:"LOG_LEVEL" default:"DEBUG"`
		Postgres    PostgresConfig
		Rpc         rpc.RpcConfig
		Web         web.WebConfig
	}
)

func InitConfig() (cfg AppConfig, err error) {
	if err = dotenv.Load(dotenv.EnvironmentFiles(os.Getenv("ENVIRONMENT"))); err != nil {
		return
	}

	err = envconfig.Process("", &cfg)

	return
}
