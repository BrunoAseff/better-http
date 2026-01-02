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
		return
	}

	defer file.Close()

	buf, err := io.ReadAll(file)

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	if len(buf) > 0 {

		for i := 0; len(buf) > i; i += 8 {

			if i > len(buf) {
				break
			}

			lo := i
			hi := i + 8

			if hi > len(buf) {
				hi = len(buf)
			}

			fmt.Printf("read: %v\n", string(buf[lo:hi]))

		}
	}
}
