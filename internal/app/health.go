package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthCheckHandler struct{}

func NewHealthCheckHandler(e *echo.Echo) *HealthCheckHandler {
	handler := &HealthCheckHandler{}
	e.GET("/", handler.HealthCheck)
	return handler
}

func (hc *HealthCheckHandler) HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}
