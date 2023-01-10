package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func HandleHealth() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(
			http.StatusOK,
			HealthResponse{Status: "ok"},
		)
	}
}
