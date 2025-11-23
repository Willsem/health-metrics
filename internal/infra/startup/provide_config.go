package startup

import (
	"health-metrics/internal/api/telegrambot"
	"health-metrics/internal/infra/config"
	"health-metrics/internal/infra/logger"
	"health-metrics/internal/infra/postgres"

	"go.uber.org/fx"
)

func provideConfig() fx.Option {
	return fx.Options(
		fx.Provide(config.New),

		fx.Provide(func(config *config.Config) *logger.Config {
			return config.Logger
		}),

		fx.Provide(func(config *config.Config) *telegrambot.Config {
			return config.Bot
		}),

		fx.Provide(func(config *config.Config) *postgres.Config {
			return config.Postgres
		}),
	)
}
