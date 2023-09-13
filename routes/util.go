package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func bindAndValidate[T any](c echo.Context, v *T) error {
	if err := c.Bind(v); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	if err := c.Validate(v); err != nil {
		return err
	}

	return nil
}
