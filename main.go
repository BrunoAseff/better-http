package main

import (
	"fmt"
	"io"
	"os"
)

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

			defer close(ch)
		}()
	}

	return ch
}

func main() {

	file, err := os.Open("messages.txt")

	if err != nil {
		fmt.Println("File could not be opened")
		return
	}

	ch := getLinesChannel(file)

	for str := range ch {
		fmt.Printf("read: %v\n", str)
	}

	defer file.Close()

}
