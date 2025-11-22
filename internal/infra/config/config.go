package config

import (
	"health-metrics/internal/api/telegrambot"
	"health-metrics/internal/infra/logger"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Logger *logger.Config      `envconfig:"LOGGER"`
	Bot    *telegrambot.Config `envconfig:"BOT"`
}

func New() (*Config, error) {
	config := &Config{}
	err := envconfig.Process("", config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
