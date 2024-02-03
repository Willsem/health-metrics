package router

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"

	"github.com/Willsem/health-metrics/internal/infra/logger"
)

type Router struct {
	echoInstance *echo.Echo
}

func New(lc fx.Lifecycle, config *Config, log logger.Logger, validator echo.Validator) *Router {
	e := echo.New()
	e.Validator = validator
	e.HideBanner = true
	e.HidePort = true

	r := &Router{
		echoInstance: e,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := r.echoInstance.Start(fmt.Sprintf(":%d", config.Port)); err != nil && err != http.ErrServerClosed {
					log.WithError(err).Fatal("failed to run the server")
				}
			}()
			return nil
		},

		OnStop: func(ctx context.Context) error {
			return r.echoInstance.Shutdown(ctx)
		},
	})

	return r
}

func (r *Router) Use(middlewares ...echo.MiddlewareFunc) {
	r.echoInstance.Use(middlewares...)
}

func (r *Router) Register(handlers ...Handler) {
	for _, h := range handlers {
		switch h.Method() {
		case http.MethodGet:
			r.echoInstance.GET(h.Path(), h.Handler())
		case http.MethodPost:
			r.echoInstance.POST(h.Path(), h.Handler())
		case http.MethodPut:
			r.echoInstance.PUT(h.Path(), h.Handler())
		case http.MethodDelete:
			r.echoInstance.DELETE(h.Path(), h.Handler())
		default:
			panic(fmt.Sprintf("unsopported method for path %s: %s", h.Path(), h.Method()))
		}
	}
}
