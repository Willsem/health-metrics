package startup

import (
	"health-metrics/internal/infra/config"
	"health-metrics/internal/infra/logger"

	"go.uber.org/fx"
)

func provideConfig() fx.Option {
	return fx.Options(
		fx.Provide(config.New),

		fx.Provide(func(config *config.Config) *logger.Config {
			return config.Logger
		}),
	)
}
