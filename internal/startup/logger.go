package startup

import (
	"go.uber.org/zap/zapcore"

	"github.com/Willsem/health-metrics/internal/infra/logger"
	"github.com/Willsem/health-metrics/internal/infra/logger/zap"
)

type LogConfig struct {
	Level zapcore.Level `envconfig:"LEVEL" default:"debug"`
	Env   string        `envconfig:"ENV" required:"true"`
}

func NewLogger(name string, config *LogConfig) logger.Logger {
	return zap.NewLogger(
		zap.Name(name),
		zap.LogLevel(config.Level),
		zap.Env(config.Env),
	)
}

func NewFallbackLogger(name string) logger.Logger {
	return zap.NewLogger(zap.Name(name))
}
