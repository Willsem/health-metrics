package router

import "github.com/labstack/echo/v4"

type Handler interface {
	Method() string
	Path() string
	Handler() echo.HandlerFunc
}
