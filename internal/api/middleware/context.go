package middleware

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	appContext "github.com/Willsem/health-metrics/internal/context"
)

type Context struct{}

func NewContext() *Context {
	return &Context{}
}

func (cm *Context) MiddlewareFunc() echo.MiddlewareFunc {
	return cm.middlewareFunc
}

func (cm *Context) middlewareFunc(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set(appContext.RequestIDKey, uuid.NewString())

		return next(c)
	}
}
