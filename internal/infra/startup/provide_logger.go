package startup

import (
	"health-metrics/internal/infra/logger"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func provideLogger() fx.Option {
	return fx.Options(
		fx.Provide(func(config *logger.Config) *logger.ZapLogger {
			return logger.NewZap(logger.GetOptions(config)...)
		}),

		fx.WithLogger(func(logger *logger.ZapLogger) fxevent.Logger {
			return &fxevent.ZapLogger{
				Logger: logger.GetSourceLogger(),
			}
		}),

		fx.Provide(func(logger *logger.ZapLogger) logger.Logger {
			return logger
		}),
	)
}
