package main

import (
	"fmt"
	"net"

	"github.com/BrunoAseff/better-http/internal/request"
)

func main() {

	const PORT = ":42069"

	listener, err := net.Listen("tcp", PORT)

	if err != nil {
		fmt.Println("Could not create connection", err)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("The connection failed", err)
		}

		if conn != nil {
			fmt.Println("The connection has been accepted")

			req, err := request.RequestFromReader(conn)

			if err == nil {
				fmt.Printf("Request line:\n- Method: %v\n- Target: %v\n- Version: %v\n", req.RequestLine.Method, req.RequestLine.RequestTarget, req.RequestLine.HttpVersion)
			}

			fmt.Println("The connection has been closed")
		}
	}

}
