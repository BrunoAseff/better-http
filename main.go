package main

import (
	"fmt"
	"io"
	"net"
)

func closeConnection(net net.Listener) {

	err := net.Close()

	if err == nil {
		fmt.Println("The connection has been closed")

		return
	}

	fmt.Println("The connection was not closed")

}

func getLinesChannel(file io.ReadCloser) <-chan string {
	buf, err := io.ReadAll(file)
	ch := make(chan string)

	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	if len(buf) > 0 {

		bytesToAdd := []byte{}

		go func() {

			defer close(ch)

			defer file.Close()

			for i := 0; i < len(buf); i += 8 {
				lo := i
				hi := i + 8

				if hi > len(buf) {
					hi = len(buf)
				}

				chars := buf[lo:hi]

				for _, char := range chars {

					if string(char) == "\n" {
						ch <- string(bytesToAdd)

						bytesToAdd = []byte{}

						continue
					}

					bytesToAdd = append(bytesToAdd, char)
				}

			}
		}()
	}

	return ch
}

func main() {

	const PORT = ":42069"

	net, err := net.Listen("tcp4", PORT)

	if err != nil {
		fmt.Println("Could not create connection", err)
	}

	defer closeConnection(net)

	for {
		conn, err := net.Accept()

		if err != nil {
			fmt.Println("The connection failed", err)
			return
		}

		if conn != nil {
			fmt.Println("The connection has been accepted")

			ch := getLinesChannel(conn)

			for str := range ch {
				fmt.Printf("read: %v\n", str)
			}
		}
	}

}
