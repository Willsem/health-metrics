package middleware

import (
	"github.com/labstack/echo/v4"

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
		return next(c)
	}
}
