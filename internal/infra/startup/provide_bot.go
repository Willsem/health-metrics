package startup

import (
	apiTelegrambot "health-metrics/internal/api/telegrambot"
	apiTelegrambotHandlersBase "health-metrics/internal/api/telegrambot/handlers/base"
	apiTelegrambotMiddlewaresAuth "health-metrics/internal/api/telegrambot/middlewares/auth"
	appGetOrgCreateUser "health-metrics/internal/application/getorcreateuser"
	"health-metrics/internal/infra/logger"
	repositoryContextUser "health-metrics/internal/repository/context/user"

	"go.uber.org/fx"
)

func provideBot() fx.Option {
	return fx.Options(
		// Bot
		fx.Provide(apiTelegrambot.New),
		fx.Invoke(func(lc fx.Lifecycle, bot *apiTelegrambot.Bot) {
			lc.Append(fx.Hook{
				OnStart: bot.Start,
				OnStop:  bot.Shutdown,
			})
		}),

		// Middlewares
		fx.Provide(
			func(usecase *appGetOrgCreateUser.GetOrCreateUserUsecase, userContext *repositoryContextUser.UserContext, logger logger.Logger) *apiTelegrambotMiddlewaresAuth.AuthMiddleware {
				return apiTelegrambotMiddlewaresAuth.New(usecase, userContext, logger)
			},
		),

		// Handlers
		fx.Provide(
			func(userContext *repositoryContextUser.UserContext, logger logger.Logger) *apiTelegrambotHandlersBase.Handler {
				return apiTelegrambotHandlersBase.New(userContext, logger)
			},
		),
	)
}
