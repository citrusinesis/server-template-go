package testing

import (
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
)

type Requester struct {
	e *echo.Echo
}

func NewRequester(echo *echo.Echo) *Requester {
	return &Requester{echo}
}

func (r *Requester) GetEcho() *echo.Echo {
	return r.e
}

func (r *Requester) BuildRequestWithOptions(method string, target string, body io.Reader, options ...func(*http.Request)) (echo.Context, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, target, body)
	rec := httptest.NewRecorder()

	for _, opt := range options {
		opt(req)
	}

	return r.e.NewContext(req, rec), rec
}

func (r *Requester) GET(target string, options ...func(*http.Request)) (echo.Context, *httptest.ResponseRecorder) {
	return r.BuildRequestWithOptions(http.MethodGet, target, nil, options...)
}

func (r *Requester) POST(target string, body io.Reader, options ...func(*http.Request)) (echo.Context, *httptest.ResponseRecorder) {
	return r.BuildRequestWithOptions(http.MethodPost, target, body, options...)
}

func (r *Requester) PUT(target string, body io.Reader, options ...func(*http.Request)) (echo.Context, *httptest.ResponseRecorder) {
	return r.BuildRequestWithOptions(http.MethodPut, target, body, options...)
}

func (r *Requester) PATCH(target string, body io.Reader, options ...func(*http.Request)) (echo.Context, *httptest.ResponseRecorder) {
	return r.BuildRequestWithOptions(http.MethodPatch, target, body, options...)
}

func (r *Requester) DELETE(target string, options ...func(*http.Request)) (echo.Context, *httptest.ResponseRecorder) {
	return r.BuildRequestWithOptions(http.MethodDelete, target, nil, options...)
}

func (r *Requester) OPTIONS(target string, options ...func(*http.Request)) (echo.Context, *httptest.ResponseRecorder) {
	return r.BuildRequestWithOptions(http.MethodOptions, target, nil, options...)
}

func (r *Requester) HEAD(target string, options ...func(*http.Request)) (echo.Context, *httptest.ResponseRecorder) {
	return r.BuildRequestWithOptions(http.MethodHead, target, nil, options...)
}

func WithJSONHeader(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
}

func WithBearerToken(token string) func(*http.Request) {
	return func(req *http.Request) {
		req.Header.Set("Authorization", "Bearer "+token)
	}
}

func WithHeader(key, value string) func(*http.Request) {
	return func(req *http.Request) {
		req.Header.Set(key, value)
	}
}
