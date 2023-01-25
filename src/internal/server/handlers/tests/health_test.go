package tests

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"url-shortener/internal/server/handlers"
)

func TestHandleHealth(t *testing.T) {
	t.Run("check health service", func(t *testing.T) {
		var (
			expectedResponse = handlers.HealthResponse{Status: "ok"}
			request          = httptest.NewRequest(http.MethodGet, "/api/health", nil)
			recorder         = httptest.NewRecorder()
			handler          = handlers.HandleHealth()
			e                = echo.New()
			c                = e.NewContext(request, recorder)
		)

		var resultResp = new(handlers.HealthResponse)
		require.NoError(t, handler(c))
		require.NoError(t, json.Unmarshal(recorder.Body.Bytes(), resultResp))
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, &expectedResponse, resultResp)
	})
}
