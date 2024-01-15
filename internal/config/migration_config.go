package config

import (
	"fmt"

	"github.com/Willsem/health-metrics/internal/infra/database"
	"github.com/kelseyhightower/envconfig"
)

type Migration struct {
	Database *database.Config `envconfig:"DB"`
}

func NewMigration() (*Migration, error) {
	config := &Migration{}

	err := envconfig.Process("", config)
	if err != nil {
		return nil, fmt.Errorf("envconfig process: %w", err)
	}

	return config, nil
}
