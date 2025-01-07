package middleware

import (
	appLog "example/pkg/log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Logger(logger appLog.Logger) echo.MiddlewareFunc {
	logger = logger.WithModule("request")
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			fields := appLog.Fields{
				"URI":    values.URI,
				"method": values.Method,
				"status": values.Status,
			}
			if values.Error != nil {
				fields["error"] = values.Error.Error()
				logger.WithFields(fields).Error("request error")
			} else {
				logger.WithFields(fields).Info("request ok")
			}
			return nil
		},
	})
}
