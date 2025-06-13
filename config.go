package server

import (
	"time"

	"github.com/numbergroup/cleanenv"
)

type Config struct {
	ReadTimeout     time.Duration `env:"READ_TIMEOUT" env-default:"5s"`
	ShutdownTimeout time.Duration `env:"SHUTDOWN_TIMEOUT" env-default:"15s"`
	Listen          string        `env:"LISTEN" env-default:":8080"`
	HealthCheckPath string        `env:"HEALTH_CHECK_PATH" env-default:"/health"`
}

func LoadServerConfigFromEnv() (Config, error) {
	conf := Config{}
	err := cleanenv.ReadEnv(&conf)

	return conf, err
}
