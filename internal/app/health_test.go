package app_test

import (
	"net/http"
	"testing"

	"example/internal/app"
	testutil "example/pkg/testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type HealthCheckSuite struct {
	suite.Suite
	requester *testutil.Requester
}

func (s *HealthCheckSuite) SetupTest() {
	s.requester = testutil.NewRequester(echo.New())
}

func (s *HealthCheckSuite) TestHealthCheckReturnsOK() {
	ctx, rec := s.requester.GET("/")
	app.HealthCheck(ctx)

	s.Equal(http.StatusOK, rec.Code)
	s.Equal("ok", rec.Body.String())
}

func TestHealthCheckSuite(t *testing.T) {
	suite.Run(t, new(HealthCheckSuite))
}
