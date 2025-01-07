package presentation

import (
	"example/internal/domain/example/application"

	"github.com/labstack/echo/v4"
)

type ExampleHandler struct {
	service *application.ExampleUsecase
}

func NewExampleHandler(e *echo.Echo, service *application.ExampleUsecase) {
	handler := &ExampleHandler{service: service}

	// Register routes
	e.GET("/example", handler.List)
}

func (h *ExampleHandler) List(c echo.Context) error {
	return c.JSON(200, map[string]string{"status": "ok"})
}
