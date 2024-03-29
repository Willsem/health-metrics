package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"

	"github.com/Willsem/health-metrics/internal/infra/database"
)

type Migration struct {
	Database *database.Config `envconfig:"DB" required:"true"`
}

func NewMigration() (*Migration, error) {
	config := &Migration{}

	err := envconfig.Process("", config)
	if err != nil {
		return nil, fmt.Errorf("envconfig process: %w", err)
	}

	return config, nil
}
