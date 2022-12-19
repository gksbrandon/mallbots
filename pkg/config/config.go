package config

import (
	"mallbots/pkg/rpc"
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
	}
)

func InitConfig() (cfg AppConfig, err error) {
	if err = dotenv.Load(dotenv.EnvironmentFiles(os.Getenv("ENVIRONMENT"))); err != nil {
		return
	}

	err = envconfig.Process("", &cfg)

	return
}
