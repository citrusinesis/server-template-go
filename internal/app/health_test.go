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
	handler   *app.HealthCheckHandler
}

func (s *HealthCheckSuite) SetupTest() {
	s.requester = testutil.NewRequester(echo.New())
	s.handler = app.NewHealthCheckHandler(s.requester.GetEcho())
}

func (s *HealthCheckSuite) TestHealthCheckReturnsOK() {
	ctx, rec := s.requester.GET("/")
	err := s.handler.HealthCheck(ctx)
	s.NoError(err)
	s.Equal(http.StatusOK, rec.Code)
	s.Equal("ok", rec.Body.String())
}

func TestHealthCheckSuite(t *testing.T) {
	suite.Run(t, new(HealthCheckSuite))
}
