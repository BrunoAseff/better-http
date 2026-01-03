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
		fmt.Println("Could resolve UDP Addr", err)
	}

	conn, err := net.DialUDP(NETWORK, nil, add)

	if err != nil {
		fmt.Println("Could not create connection", err)
	}

	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("> ")

		str, err := reader.ReadString('\n')

		if err != nil {
			fmt.Printf("An error occurred while reading the string", err)
		}

		_, err = conn.Write([]byte(str))

		if err != nil {
			fmt.Printf("An error occurred while writing in the connection", err)
		}

	}
}
