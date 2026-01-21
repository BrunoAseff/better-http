package headers

import (
	"bytes"
	"errors"
	"strings"
)

type Headers map[string]string

func NewHeaders() Headers {
	return make(map[string]string)
}

func (h Headers) Parse(data []byte) (int, bool, error) {
	totalRead := 0

	for {
		currentData := data[totalRead:]
		idx := bytes.Index(currentData, []byte("\r\n"))

		if idx == -1 {
			return totalRead, false, nil
		}

		if idx == 0 {
			return totalRead + 2, true, nil
		}

		line := currentData[:idx]
		colonIdx := bytes.IndexByte(line, ':')

		if colonIdx == -1 {
			return 0, false, errors.New("malformed header: no colon found")
		}

		keyBytes := line[:colonIdx]
		valBytes := line[colonIdx+1:]

		if len(keyBytes) > 0 && keyBytes[len(keyBytes)-1] == ' ' {
			return 0, false, errors.New("header field name cannot have trailing whitespace before colon")
		}

		key := strings.ToLower(string(keyBytes))

		val := string(bytes.TrimSpace(valBytes))

		if !isAllowed(key) {
			return 0, false, errors.New("malformed header: invalid characters in key")
		}

		if key == "" {
			return 0, false, errors.New("malformed header: empty key")
		}

		h[key] = val

		totalRead += idx + 2
	}
}

func isAllowed(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, r := range s {
		if !isAllowedRune(r) {
			return false
		}
	}
	return true
}

func isAllowedRune(r rune) bool {
	if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
		return true
	}
	return strings.ContainsRune("!#$%&'*+-.^_`|~", r)
}
