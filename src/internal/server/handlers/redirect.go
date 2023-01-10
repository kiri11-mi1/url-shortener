package handlers

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"url-shortener/internal/model"
)

type redirecter interface {
	Redirect(ctx context.Context, identifier string) (string, error)
}

func HandleRedirect(redirecter redirecter) echo.HandlerFunc {
	return func(c echo.Context) error {
		identifier := c.Param("identifier")
		redirectURL, err := redirecter.Redirect(c.Request().Context(), identifier)
		if err != nil {
			if errors.Is(err, model.ErrNotFound) {
				return echo.NewHTTPError(http.StatusNotFound, err.Error())
			}
			log.Printf("redirect failed: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.Redirect(http.StatusMovedPermanently, redirectURL)
	}
}
