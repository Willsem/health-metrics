package middleware

import (
	"github.com/labstack/echo/v4"

	appContext "github.com/Willsem/health-metrics/internal/context"
	"github.com/Willsem/health-metrics/internal/infra/logger"
)

type Logger struct {
	log logger.Logger
}

func NewLogger(log logger.Logger) *Logger {
	return &Logger{
		log: log,
	}
}

func (l *Logger) MiddlewareFunc() echo.MiddlewareFunc {
	return l.middlewareFunc
}

func (l *Logger) middlewareFunc(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log := l.log.
			With(appContext.RequestIDKey, c.Get(appContext.RequestIDKey).(string)).
			With("path", c.Path()).
			With("method", c.Request().Method)

		log.Info("request started")

		err := next(c)
		if err != nil {
			log.WithError(err).Warn("error during the request")
			return err
		}

		log.Info("request completed successfully")
		return nil
	}
}
