package app_test

import (
	"net/http"
	"testing"

	"example/internal/app"
	testutil "example/pkg/testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	t.Parallel()

	t.Run("should return status 200", func(t *testing.T) {
		ctx, rec := testutil.GET("/")
		app.HealthCheck(ctx)
		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("should return string \"ok\"", func(t *testing.T) {
		ctx, rec := testutil.GET("/")
		app.HealthCheck(ctx)
		assert.Equal(t, "ok", rec.Body.String())
	})
}
