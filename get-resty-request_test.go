package core_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"maschine.io/core"
)

func TestGetRestyRequest(t *testing.T) {
	t.Run("request initialization", func(t *testing.T) {
		req := core.GetRestyRequest()
		assert.NotNil(t, req, "Request should not be nil")
	})

	t.Run("header validation", func(t *testing.T) {
		req := core.GetRestyRequest()
		headers := req.Header

		expectedHeaders := map[string]string{
			"User-Agent":   "maschine/1.0",
			"Content-Type": "application/json",
			"Accept":       "application/json",
		}

		for key, expectedValue := range expectedHeaders {
			value := headers.Get(key)
			assert.Equal(t, expectedValue, value, "Header %s should be %s, got %s", key, expectedValue, value)
		}
	})

	t.Run("header count", func(t *testing.T) {
		req := core.GetRestyRequest()
		headers := req.Header
		assert.Equal(t, 3, len(headers), "Should have exactly 3 headers set")
	})

	t.Run("individual headers", func(t *testing.T) {
		req := core.GetRestyRequest()
		headers := req.Header

		// Test User-Agent header
		assert.Equal(t, "maschine/1.0", headers.Get("User-Agent"))

		// Test Content-Type header
		assert.Equal(t, "application/json", headers.Get("Content-Type"))

		// Test Accept header
		assert.Equal(t, "application/json", headers.Get("Accept"))
	})
}
