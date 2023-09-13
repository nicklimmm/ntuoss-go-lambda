package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// registerRoot registers routes under "/"
func registerRoot(e *echo.Echo) {
	e.GET("/healthcheck", healthcheck)
}

func healthcheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
