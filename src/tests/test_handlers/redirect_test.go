package test_handlers

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/samber/mo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"url-shortener/internal/model"
	"url-shortener/internal/server/handlers"
	"url-shortener/internal/shorten"
	"url-shortener/internal/storage/shortening"
)

func TestHandleRedirect(t *testing.T) {
	t.Run("redirect to original URL", func(t *testing.T) {
		const (
			originalURL = "https://example.com"
			identifier  = "example"
		)
		var (
			redirecter = shorten.NewService(shortening.NewInMemory())
			request    = httptest.NewRequest(http.MethodGet, "/"+identifier, nil)
			recorder   = httptest.NewRecorder()
			handler    = handlers.HandleRedirect(redirecter)
			e          = echo.New()
			c          = e.NewContext(request, recorder)
		)
		c.SetPath("/:identifier")
		c.SetParamNames("identifier")
		c.SetParamValues(identifier)

		_, err := redirecter.Shorten(
			context.Background(),
			model.ShortenInput{
				RawURL:     originalURL,
				Identifier: mo.Some(identifier),
			},
		)
		require.NoError(t, err)
		require.NoError(t, handler(c))
		assert.Equal(t, http.StatusMovedPermanently, recorder.Code)
		assert.Equal(t, originalURL, recorder.Header().Get("Location"))
	})

	t.Run("not found identifier", func(t *testing.T) {
		const identifier = "testExample"
		var (
			redirecter = shorten.NewService(shortening.NewInMemory())
			request    = httptest.NewRequest(http.MethodGet, "/"+identifier, nil)
			recorder   = httptest.NewRecorder()
			handler    = handlers.HandleRedirect(redirecter)
			e          = echo.New()
			c          = e.NewContext(request, recorder)
		)

		c.SetPath("/:identifier")
		c.SetParamNames("identifier")
		c.SetParamValues(identifier)

		require.Error(t, handler(c))
	})
}
