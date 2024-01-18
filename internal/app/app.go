package app

import (
	"context"
	"fmt"

	"github.com/Willsem/health-metrics/internal/config"
	"github.com/Willsem/health-metrics/internal/startup"
)

func Run(ctx context.Context, appName string) error {
	cfg, err := config.NewApp()
	if err != nil {
		return fmt.Errorf("parse config: %w", err)
	}

	_ = startup.NewLogger(appName, cfg.Log)

	return nil
}
