package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
)

func main() {

	const PORT = ":42069"

	listener, err := net.Listen("tcp", PORT)

	if err != nil {
		fmt.Println("Could not create connection\n", err)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("The connection failed\n", err)
		}

		if conn != nil {
			fmt.Println("The connection has been accepted\n")

			ch := getLinesChannel(conn)

			for str := range ch {
				fmt.Println(str)
			}

			fmt.Println("The connection has been closed")
		}
	}

}

func getLinesChannel(f io.ReadCloser) <-chan string {

	lines := make(chan string)

	go func() {
		defer f.Close()
		defer close(lines)

		currentLineContents := ""

		for {

			b := make([]byte, 8, 8)

			n, err := f.Read(b)

			if err != nil {
				if currentLineContents != "" {
					lines <- currentLineContents
				}
				if errors.Is(err, io.EOF) {
					break
				}
				fmt.Printf("error: %s\n", err.Error())
				return
			}

			str := string(b[:n])

			parts := strings.Split(str, "\n")

			for i := 0; i < len(parts)-1; i++ {
				lines <- fmt.Sprintf("%s%s", currentLineContents, parts[i])
				currentLineContents = ""
			}

			currentLineContents += parts[len(parts)-1]

		}
	}()

	return lines
}
