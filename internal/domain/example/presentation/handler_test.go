package presentation_test

import (
	"net/http"
	"testing"

	"example/internal/domain/example/presentation"
	testutil "example/pkg/testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type ExampleHandlerSuite struct {
	suite.Suite
	requester *testutil.Requester
	handler   *presentation.ExampleHandler
}

func (s *ExampleHandlerSuite) SetupTest() {
	s.requester = testutil.NewRequester(echo.New())
	s.handler = presentation.NewExampleHandler(s.requester.GetEcho(), nil)
}

func (s *ExampleHandlerSuite) TestList() {
	ctx, rec := s.requester.GET("/example")
	err := s.handler.List(ctx)
	s.NoError(err)
	s.JSONEq(`{"test":"ok"}`, rec.Body.String())
	s.Equal(http.StatusOK, rec.Code)
}

func TestExampleHandlerSuite(t *testing.T) {
	suite.Run(t, new(ExampleHandlerSuite))
}
