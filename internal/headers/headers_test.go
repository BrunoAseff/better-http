package headers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHeaders(t *testing.T) {

	t.Run("Valid single header", func(t *testing.T) {
		headers := NewHeaders()
		data := []byte("Host: localhost:42069\r\n")
		n, done, err := headers.Parse(data)

		require.NoError(t, err)
		assert.Equal(t, "localhost:42069", headers["Host"])
		assert.Equal(t, 23, n)
		assert.False(t, done)
	})

	t.Run("Valid single header with extra whitespace", func(t *testing.T) {
		// HTTP spec requires trimming whitespace around the value
		headers := NewHeaders()
		data := []byte("Content-Type:   application/json    \r\n")
		n, done, err := headers.Parse(data)

		require.NoError(t, err)
		assert.Equal(t, "application/json", headers["Content-Type"])
		assert.Equal(t, len(data), n)
		assert.False(t, done)
	})

	t.Run("Valid 2 headers with existing headers", func(t *testing.T) {
		headers := NewHeaders()
		// First pass
		_, _, _ = headers.Parse([]byte("User-Agent: Go-Test\r\n"))

		// Second pass adding more headers
		data := []byte("Accept: */*\r\nConnection: keep-alive\r\n")
		n, done, err := headers.Parse(data)

		require.NoError(t, err)
		assert.Equal(t, "Go-Test", headers["User-Agent"])
		assert.Equal(t, "*/*", headers["Accept"])
		assert.Equal(t, "keep-alive", headers["Connection"])
		assert.Equal(t, len(data), n)
		assert.False(t, done)
	})

	t.Run("Valid done", func(t *testing.T) {
		// An empty line (\r\n) after headers signals the end of the header block
		headers := NewHeaders()
		data := []byte("\r\n")
		n, done, err := headers.Parse(data)

		require.NoError(t, err)
		assert.Equal(t, 2, n)
		assert.True(t, done)
	})

	t.Run("Invalid spacing header", func(t *testing.T) {
		// Spaces before the colon are technically illegal in many strict HTTP parsers (RFC 7230)
		headers := NewHeaders()
		data := []byte("Host : localhost:42069\r\n")
		n, done, err := headers.Parse(data)

		require.Error(t, err)
		assert.Equal(t, 0, n)
		assert.False(t, done)
	})
}
