package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"url-shortener/internal/server/handlers"
)

func TestHandleHealth(t *testing.T) {
	t.Run("check health service", func(t *testing.T) {
		var expectedResponse = handlers.HealthResponse{Status: "ok"}
		var request = httptest.NewRequest(http.MethodGet, "/api/health", nil)
	})
}
