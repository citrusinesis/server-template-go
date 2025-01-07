package presentation

import (
	"example/internal/domain/example/application"

	"github.com/labstack/echo/v4"
)

type ExampleHandler struct {
	service *application.ExampleUsecase
}

func NewExampleHandler(e *echo.Echo, service *application.ExampleUsecase) *ExampleHandler {
	handler := &ExampleHandler{service: service}

	// Register routes
	e.GET("/example", handler.List)

	return handler
}

func (h *ExampleHandler) List(c echo.Context) error {
	return c.JSON(200, map[string]string{"test": "ok"})
}
