package startup

import (
	"health-metrics/internal/api/telegrambot"

	"go.uber.org/fx"
)

func provideBot() fx.Option {
	return fx.Options(
		fx.Provide(telegrambot.New),

		fx.Invoke(func(lc fx.Lifecycle, bot *telegrambot.Bot) {
			lc.Append(fx.Hook{
				OnStart: bot.Start,
				OnStop:  bot.Shutdown,
			})
		}),
	)
}
