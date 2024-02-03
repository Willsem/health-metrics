package app

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"

	"github.com/Willsem/health-metrics/internal/api/handlers"
	"github.com/Willsem/health-metrics/internal/api/router"
	"github.com/Willsem/health-metrics/internal/config"
	"github.com/Willsem/health-metrics/internal/infra/logger"
	"github.com/Willsem/health-metrics/internal/infra/logger/zap"
	"github.com/Willsem/health-metrics/internal/startup"
)

func WithLogger() fx.Option {
	return fx.WithLogger(func(log logger.Logger) fxevent.Logger {
		zapLogger := log.(*zap.LoggerImpl)
		return &fxevent.ZapLogger{Logger: zapLogger.GetZapLogger().Desugar()}
	})
}

func ProvideConfig() fx.Option {
	cfg, err := config.NewApp()
	if err != nil {
		return fx.Error(err)
	}

	return fx.Provide(func() *config.App {
		return cfg
	})
}

func ProvideLogger(appName string) fx.Option {
	return fx.Provide(func(cfg *config.App) logger.Logger {
		return startup.NewLogger(appName, cfg.Log)
	})
}

func ProvideRouter() fx.Option {
	return fx.Provide(
		router.NewValidator,

		func(lc fx.Lifecycle, cfg *config.App, log logger.Logger, validator echo.Validator) *router.Router {
			return router.New(lc, cfg.Router, log, validator)
		},
	)
}

func ProvideHandlers() fx.Option {
	return fx.Provide(
		handlers.NewPostLogin,
	)
}

func RegisterHandlers() fx.Option {
	return fx.Invoke(func(router *router.Router, postLogin *handlers.PostLogin) {
		router.Register(postLogin)
	})
}
