package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"

	"github.com/Willsem/health-metrics/internal/infra/database"
	"github.com/Willsem/health-metrics/internal/startup"
)

type App struct {
	Database *database.Config   `envconfig:"DB" required:"true"`
	Log      *startup.LogConfig `envconfig:"LOG" required:"true"`
}

func NewApp() (*App, error) {
	config := &App{}

	err := envconfig.Process("", config)
	if err != nil {
		return nil, fmt.Errorf("envconfig process: %w", err)
	}

	return config, nil
}
