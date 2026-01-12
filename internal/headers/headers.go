package headers

import (
	"bytes"
	"errors"
)

type Headers map[string]string

func NewHeaders() Headers {
	headers := Headers{}

	return headers
}

func (h Headers) Parse(data []byte) (n int, done bool, err error) {

	idx := bytes.Index(data, []byte("\r\n"))

	if idx == -1 {
		return 0, false, nil
	}
	parts := bytes.Split(data, []byte(":"))

	if len(parts) != 2 {

		err := errors.New("Incorect header format")

		return 0, false, err
	}

	key, value := parts[0], parts[1]

	lastChar := len(string(key)) - 1

	if string(lastChar) == " " {
		err := errors.New("Incorect header format")

		return 0, true, err
	}

	h[string(key)] = string(value)

	n = len(key) + len(value)

	return n, true, nil
}
