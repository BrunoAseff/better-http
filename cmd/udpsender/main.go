package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	const ADDRESS = ":42069"
	const NETWORK = "udp"

	add, err := net.ResolveUDPAddr(NETWORK, ADDRESS)

	if err != nil {
		fmt.Printf("Could resolve UDP Addr: %v", err)
	}

	conn, err := net.DialUDP(NETWORK, nil, add)

	if err != nil {
		fmt.Printf("Could not create connection: %v", err)
	}

	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("> ")

		str, err := reader.ReadString('\n')

		if err != nil {
			fmt.Printf("An error occurred while reading the string: %v", err)
		}

		_, err = conn.Write([]byte(str))

		if err != nil {
			fmt.Printf("An error occurred while writing in the connection: %v", err)
		}

	}
}
