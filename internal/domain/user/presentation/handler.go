package presentation

import (
	"example/internal/domain/user/application"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service *application.UserExampleUsecase
}

func NewUserHandler(e *echo.Echo, service *application.UserExampleUsecase) {
	handler := &UserHandler{service: service}

	// Register routes
	e.GET("/users", handler.List)
}

func (h *UserHandler) List(c echo.Context) error {
	return c.JSON(200, map[string]string{"status": "ok"})
}
