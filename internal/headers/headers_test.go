package headers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHeaders(t *testing.T) {

	t.Run("Valid single header - Case Insensitivity", func(t *testing.T) {
		headers := NewHeaders()
		data := []byte("hOsT: localhost:42069\r\n")
		n, done, err := headers.Parse(data)

		require.NoError(t, err)
		assert.Equal(t, "localhost:42069", headers["host"])
		assert.Equal(t, 23, n)
		assert.False(t, done)
	})

	t.Run("Valid single header with extra whitespace", func(t *testing.T) {
		headers := NewHeaders()
		data := []byte("Content-Type:   application/json    \r\n")
		n, done, err := headers.Parse(data)

		require.NoError(t, err)
		assert.Equal(t, "application/json", headers["content-type"])
		assert.Equal(t, len(data), n)
		assert.False(t, done)
	})

	t.Run("Valid 2 headers with existing headers", func(t *testing.T) {
		headers := NewHeaders()
		_, _, _ = headers.Parse([]byte("User-Agent: Go-Test\r\n"))

		data := []byte("Accept: */*\r\nConnection: keep-alive\r\n")
		n, done, err := headers.Parse(data)

		require.NoError(t, err)
		assert.Equal(t, "Go-Test", headers["user-agent"])
		assert.Equal(t, "*/*", headers["accept"])
		assert.Equal(t, "keep-alive", headers["connection"])
		assert.Equal(t, len(data), n)
		assert.False(t, done)
	})

	t.Run("Invalid character in key", func(t *testing.T) {
		headers := NewHeaders()
		data := []byte("H©st: localhost:42069\r\n")
		n, done, err := headers.Parse(data)

		require.Error(t, err, "Should error on invalid character ©")
		assert.Equal(t, 0, n)
		assert.False(t, done)
	})

	t.Run("Valid done", func(t *testing.T) {
		headers := NewHeaders()
		data := []byte("\r\n")
		n, done, err := headers.Parse(data)

		require.NoError(t, err)
		assert.Equal(t, 2, n)
		assert.True(t, done)
	})
}
