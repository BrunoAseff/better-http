package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Open("messages.txt")

	if err != nil {
		fmt.Println("File could not be opened")
	}

	defer file.Close()

	buf := make([]byte, 1024)
	var num int

	for {
		n, err := file.Read(buf)

		num += n

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("Error reading the file")
			break
		}
	}

	if len(buf) > 0 {

		for i := 0; len(buf) > i; i += 8 {

			if i > num {
				break
			}

			lo := i
			hi := i + 8

			fmt.Printf("read: %v\n", string(buf[lo:hi]))

		}
	}
}
