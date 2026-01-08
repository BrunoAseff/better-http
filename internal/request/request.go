package request

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
)

type parserState int

const (
	initialized parserState = iota
	done
)

type Request struct {
	RequestLine RequestLine
	state       parserState
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func (r *Request) parse(data []byte) (int, error) {
	if r.state == initialized {
		requestLine, bytesNum, err := parseRequestLine(data)

		if err != nil {
			return 0, err
		}

		if bytesNum == 0 {
			return 0, nil
		}

		r.RequestLine = requestLine
		r.state = done

		return bytesNum, nil
	}

	if r.state == done {
		return 0, errors.New("error: trying to read data in a done state")
	}

	return 0, errors.New("error: unknown state")
}

func RequestFromReader(reader io.Reader) (*Request, error) {

	data, err := io.ReadAll(reader)

	if err != nil {
		err = errors.New("error while reading the string")

		return nil, err
	}

	requestLine, _, err := parseRequestLine(data)

	if err != nil {
		return nil, err
	}

	return &Request{
		RequestLine: requestLine,
	}, nil

}

func parseRequestLine(request []byte) (rl RequestLine, bytesNum int, err error) {

	separator := []byte{'\r', '\n'}

	lines := bytes.Split(request, separator)

	if len(lines) < 2 {
		return RequestLine{}, 0, nil
	}

	requestLine := lines[0]

	bytesNum = len(requestLine) + len(separator)

	sections := strings.Split(string(requestLine), " ")

	if len(sections) != 3 {
		err := errors.New("the request line is in the incorrect format")

		return RequestLine{}, 0, err
	}

	httpSections := strings.Split(sections[2], "/")

	if len(httpSections) != 2 {
		err := errors.New("invalid HTTP version")

		return RequestLine{}, 0, err
	}

	method := sections[0]
	requestTarget := sections[1]
	httpVersion := httpSections[1]

	if method == "" {
		err := errors.New("method was not provided in the request line")

		return RequestLine{}, 0, err
	}

	for _, c := range method {
		if c < 'A' || c > 'Z' {
			return RequestLine{}, 0, fmt.Errorf("invalid method: %s", method)
		}
	}

	if httpVersion != "1.1" {

		err := fmt.Errorf("invalid HTTP version\nExpected: HTTP/1.1\nReceived: %v", httpVersion)

		return RequestLine{}, 0, err
	}

	return RequestLine{
		Method:        method,
		HttpVersion:   httpVersion,
		RequestTarget: requestTarget,
	}, bytesNum, nil
}
