package shorten_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"url-shortener/internal/shorten"
)

func TestShorten(t *testing.T) {
	t.Run("return alpha-numeric ids", func(t *testing.T) {
		type testCase struct {
			id       uint32
			expected string
		}
		testCases := []testCase{
			{id: 0, expected: ""},
			{id: 512, expected: "oM"},
		}
		for _, tc := range testCases {
			result := shorten.Shorten(tc.id)
			assert.Equal(t, tc.expected, result)
		}
	})
	t.Run("is idempotent ", func(t *testing.T) {
		type testCase struct {
			id       uint32
			expected string
		}
		tc := testCase{839, "gE"}
		for i := 0; i < 100; i++ {
			assert.Equal(t, tc.expected, shorten.Shorten(tc.id))
		}
	})
}

func TestService_PrependBaseURL(t *testing.T) {
	t.Run("get new url", func(t *testing.T) {
		var (
			baseURL    = "https://example.com"
			identifier = "testId83"
			expected   = "https://example.com/testId83"
		)
		var prependURL, err = shorten.PrependBaseURL(baseURL, identifier)
		require.NoError(t, err)
		assert.Equal(t, expected, prependURL)
	})
	t.Run("get new url", func(t *testing.T) {
		var (
			baseURL    = "https://example.com/kajjsndkajdasdna/oansdkjansjdk"
			identifier = "testId83"
			expected   = "https://example.com/testId83"
		)
		var prependURL, err = shorten.PrependBaseURL(baseURL, identifier)
		require.NoError(t, err)
		assert.Equal(t, expected, prependURL)
	})
}
