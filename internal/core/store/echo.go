package store

import (
	"context"
	"example/internal/core/middleware"
	"example/pkg/store"

	"github.com/labstack/echo/v4"
)

func GetFromEcho[T any](c echo.Context, key store.Key) (T, bool) {
	var defaultValue T

	ctx, ok := c.Get(middleware.ContextKey).(context.Context)
	if !ok {
		return defaultValue, ok
	}

	return Get[T](ctx, key)
}

func SetToEcho[T any](c echo.Context, key store.Key, val T) bool {
	ctx, ok := c.Get(middleware.ContextKey).(context.Context)
	if !ok {
		return ok
	}

	return Set[T](ctx, key, val)
}
