package tests

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"url-shortener/internal/server"
	"url-shortener/internal/server/handlers"
	"url-shortener/internal/shorten"
	"url-shortener/internal/storage/shortening"
)

func TestHandleShorten(t *testing.T) {
	t.Run("shorten url success", func(t *testing.T) {
		const testPayload = `{"url": "https://google.com"}`

		var (
			shortener = shorten.NewService(shortening.NewInMemory())
			handler   = handlers.HandleShorten(shortener)
			recorder  = httptest.NewRecorder()
			request   = httptest.NewRequest(http.MethodPost, "/api/shorten", strings.NewReader(testPayload))
			e         = echo.New()
			c         = e.NewContext(request, recorder)
		)
		e.Validator = server.NewValidator()
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		require.NoError(t, handler(c))
		assert.Equal(t, http.StatusOK, recorder.Code)

		var resp = handlers.ShortenResponse{}
		require.NoError(t, json.Unmarshal(recorder.Body.Bytes(), &resp))
		assert.NotEmpty(t, resp.ShortURL)
	})

	t.Run("invalid url", func(t *testing.T) {
		const testPayload = `{"url": "bla bla bla"}`

		var (
			shortener = shorten.NewService(shortening.NewInMemory())
			handler   = handlers.HandleShorten(shortener)
			recorder  = httptest.NewRecorder()
			request   = httptest.NewRequest(http.MethodPost, "/api/shorten", strings.NewReader(testPayload))
			e         = echo.New()
			c         = e.NewContext(request, recorder)
		)
		e.Validator = server.NewValidator()
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		var httpErr *echo.HTTPError
		require.ErrorAs(t, handler(c), &httpErr)
		assert.Equal(t, http.StatusBadRequest, httpErr.Code)
		assert.Contains(t, httpErr.Message, "Field validation for 'URL' failed")
	})

	t.Run("identifer already exist", func(t *testing.T) {
		const testPayload = `{"url": "https://www.google.com", "identifier": "example"}`

		var (
			shortener = shorten.NewService(shortening.NewInMemory())
			handler   = handlers.HandleShorten(shortener)
			recorder  = httptest.NewRecorder()
			request   = httptest.NewRequest(http.MethodPost, "/api/shorten", strings.NewReader(testPayload))
			e         = echo.New()
			c         = e.NewContext(request, recorder)
		)

		e.Validator = server.NewValidator()
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		require.NoError(t, handler(c))
		assert.Equal(t, http.StatusOK, recorder.Code)

		request = httptest.NewRequest(http.MethodPost, "/api/shorten", strings.NewReader(testPayload))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder = httptest.NewRecorder()
		c = e.NewContext(request, recorder)

		var httpErr *echo.HTTPError
		require.ErrorAs(t, handler(c), &httpErr)
		assert.Equal(t, http.StatusConflict, httpErr.Code)
	})
}
