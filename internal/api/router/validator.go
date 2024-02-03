package router

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Validator struct {
	validator *validator.Validate
}

func NewValidator() echo.Validator {
	return &Validator{
		validator: validator.New(),
	}
}

func (v *Validator) Validate(i any) error {
	if err := v.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}
