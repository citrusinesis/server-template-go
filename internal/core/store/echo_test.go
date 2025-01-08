package store_test

import (
	"net/http"
	"testing"

	"example/internal/core/middleware"
	store "example/internal/core/store"
	testutil "example/pkg/testing"

	"github.com/labstack/echo/v4"
)

func TestGetFromEchoAndSetToEcho(t *testing.T) {
	e := echo.New()
	e.Use(middleware.ContextStore())

	e.GET("/test", func(c echo.Context) error {
		ok := store.SetToEcho(c, store.ExampleKey, "echo value")
		if !ok {
			return c.String(http.StatusInternalServerError, "SetToEcho failed")
		}

		val, ok := store.GetFromEcho[string](c, store.ExampleKey)
		if !ok {
			return c.String(http.StatusInternalServerError, "GetFromEcho failed")
		}

		if val != "echo value" {
			return c.String(http.StatusInternalServerError, "invalid value from store")
		}

		return c.NoContent(http.StatusOK)
	})

	r := testutil.NewRequester(e)
	ctx, rec := r.GET("/test")

	e.ServeHTTP(rec, ctx.Request())

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d; body=%q", rec.Code, rec.Body.String())
	}
}
