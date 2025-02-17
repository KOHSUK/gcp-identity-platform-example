package config

import (
	"app/internal/web"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type (
	PGConfig struct {
		Conn string `required:"true"`
	}
	AppConfig struct {
		Environment     string
		LogLevel        string `envconfig:"LOG_LEVEL" default:"DEBUG"`
		PG              PGConfig
		Web             web.WebConfig
		ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" default:"30s"`
	}
)

func InitConfig() (cfg AppConfig, err error) {
	if err = godotenv.Load(); err != nil {
		return
	}

	err = envconfig.Process("", &cfg)

	return
}
