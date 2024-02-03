package handlers

import (
	"net/http"

	"github.com/Willsem/health-metrics/internal/api/dto"
	"github.com/labstack/echo/v4"
)

// PostLogin вход в систему
type PostLogin struct{}

func NewPostLogin() *PostLogin {
	return &PostLogin{}
}

func (*PostLogin) Method() string {
	return http.MethodPost
}

func (*PostLogin) Path() string {
	return "/api/v1/login"
}

func (h *PostLogin) Handler() echo.HandlerFunc {
	return h.handlerFunc
}

func (h *PostLogin) handlerFunc(c echo.Context) error {
	postUser := &dto.PostUser{}
	if err := c.Bind(postUser); err != nil {
		return err
	}

	if err := c.Validate(postUser); err != nil {
		return err
	}

	return c.JSON(200, postUser)
}
