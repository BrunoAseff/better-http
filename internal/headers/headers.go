package headers

import (
	"bytes"
	"errors"
	"strings"
	"unicode"
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

		key := string(bytes.TrimSpace(keyBytes))
		val := string(bytes.TrimSpace(valBytes))

		if !isAllowed(key) {
			return 0, false, errors.New("malformed header: invalid characters in key")
		}

		h[key] = val

		totalRead += idx + 2
	}
}

func isAllowed(s string) bool {
	return strings.IndexFunc(s, func(r rune) bool {
		return !isAllowedRune(r)
	}) == -1
}

func isAllowedRune(r rune) bool {
	if unicode.IsLetter(r) || unicode.IsDigit(r) {
		return true
	}
	return strings.ContainsRune("!#$%&'*+-.^_`|~", r)
}
