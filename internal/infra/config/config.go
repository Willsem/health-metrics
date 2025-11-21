package config

import (
	"health-metrics/internal/infra/logger"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Logger *logger.Config `envconfig:"LOGGER"`
}

func New() (*Config, error) {
	config := &Config{}
	err := envconfig.Process("", config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
