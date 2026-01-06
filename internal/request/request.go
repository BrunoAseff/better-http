package request

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"unicode/utf8"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func RequestFromReader(reader io.Reader) (*Request, error) {

	data, err := io.ReadAll(reader)

	if err != nil {
		err = errors.New("Error while reading the string")

		return nil, err
	}

	str := strings.Split(string(data), "\r\n")

	requestLine, err := parseRequestLine(str[0])

	if err != nil {
		return nil, err
	}

	return &Request{
		RequestLine: requestLine,
	}, nil

}

func parseRequestLine(line string) (RequestLine, error) {

	sections := strings.Split(line, " ")

	if len(sections) != 3 {
		err := errors.New("The request line is in the incorrect format")

		return RequestLine{}, err
	}

	method := sections[0]
	requestTarget := sections[1]
	httpVersion := strings.Split(sections[2], "/")[1]

	if method == "" {
		err := errors.New("Method was not provided in the request line")

		return RequestLine{}, err
	}

	if method != strings.ToUpper(method) || !utf8.ValidString(method) {
		err := errors.New("Method is in the incorrect format")

		return RequestLine{}, err
	}

	if httpVersion != "1.1" {

		err := errors.New(fmt.Sprintf("Invalid HTTP version\nExpected: HTTP/1.1\nReceived: %v", httpVersion))

		return RequestLine{}, err
	}

	return RequestLine{
		Method:        method,
		HttpVersion:   httpVersion,
		RequestTarget: requestTarget,
	}, nil
}
