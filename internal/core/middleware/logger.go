package middleware

import (
	appLog "example/pkg/log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Logger(logger appLog.Logger) echo.MiddlewareFunc {
	logger = logger.WithModule("request")
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:     true,
		LogStatus:  true,
		LogMethod:  true,
		LogLatency: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			fields := appLog.Fields{
				"method":  values.Method,
				"uri":     values.URI,
				"status":  values.Status,
				"latency": values.Latency,
			}
			if values.Error != nil {
				fields["error"] = values.Error.Error()
				logger.WithFields(fields).Error(values.Error.Error())
			} else {
				logger.WithFields(fields).Info()
			}
			return nil
		},
	})
}
