package tests_test

import (
	"context"
	"github.com/samber/mo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"url-shortener/internal/model"
	"url-shortener/internal/shorten"
	"url-shortener/internal/storage/shortening"
)

func TestService_Shorten(t *testing.T) {
	t.Run("generate shortening for a given url", func(t *testing.T) {
		var (
			svc   = shorten.NewService(shortening.NewInMemory())
			input = model.ShortenInput{RawURL: "https://example.com"}
		)
		createdShortening, err := svc.Shorten(context.Background(), input)
		require.NoError(t, err)

		require.NotEmpty(t, createdShortening.Identifier)
		assert.Equal(t, input.RawURL, createdShortening.OriginalUrl)
		assert.NotZero(t, createdShortening.CreatedAt)
	})
	t.Run("use custom id", func(t *testing.T) {
		const identifier = "exmpl"
		var (
			svc   = shorten.NewService(shortening.NewInMemory())
			input = model.ShortenInput{RawURL: "https://example.com", Identifier: mo.Some(identifier)}
		)
		createdShortening, err := svc.Shorten(context.Background(), input)
		require.NoError(t, err)

		assert.Equal(t, identifier, createdShortening.Identifier)
		assert.Equal(t, input.RawURL, createdShortening.OriginalUrl)
		assert.NotZero(t, createdShortening.CreatedAt)
	})
	t.Run("returns error if id is already taken", func(t *testing.T) {
		const identifier = "exmpl"
		var (
			svc   = shorten.NewService(shortening.NewInMemory())
			input = model.ShortenInput{RawURL: "https://example.com", Identifier: mo.Some(identifier)}
		)
		_, err := svc.Shorten(context.Background(), input)
		require.NoError(t, err)

		_, err = svc.Shorten(context.Background(), input)
		assert.Equal(t, err, model.ErrIdentifierIsExist)
	})
}
