package middleware

import (
	"context"

	"example/pkg/store"

	"github.com/labstack/echo/v4"
)

const ContextKey = "ctx"

type InitStore = func(s *store.Store)

func ContextStore(init ...InitStore) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			s := store.NewStore()
			for _, invokeInit := range init {
				invokeInit(s)
			}

			c.Set(
				ContextKey,
				store.WithStore(context.Background(), s),
			)
			return next(c)
		}
	}
}
