package testing

import (
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
)

func BuildRequest(method string, target string, body io.Reader) (echo.Context, *httptest.ResponseRecorder) {
	return BuildRequestWithOptions(method, target, body, nil)
}

func BuildRequestWithOptions(method string, target string, body io.Reader, options func(*http.Request)) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(method, target, body)
	rec := httptest.NewRecorder()

	if options != nil {
		options(req)
	}

	return e.NewContext(req, rec), rec
}

func GET(target string) (echo.Context, *httptest.ResponseRecorder) {
	return BuildRequest(http.MethodGet, target, nil)
}

func POST(target string, body io.Reader) (echo.Context, *httptest.ResponseRecorder) {
	return BuildRequest(http.MethodGet, target, body)
}

func PATCH(target string, body io.Reader) (echo.Context, *httptest.ResponseRecorder) {
	return BuildRequest(http.MethodPatch, target, body)
}

func DELETE(target string) (echo.Context, *httptest.ResponseRecorder) {
	return BuildRequest(http.MethodDelete, target, nil)
}
