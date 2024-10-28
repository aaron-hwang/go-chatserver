package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func client(address string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Printf("There was an error while trying to connect to the server.")
		os.Exit(1)
	}

	handleConnection(conn)
}

func handleConnection(conn net.Conn) {
	input := readFromUser()
	for input != "quit" {
		conn.Write([]byte(input))
		reader := bufio.NewReader(conn)
		output, err := io.ReadAll(reader)
		if err != nil {
			fmt.Printf("Could not read response from server")
			conn.Close()
			return
		}
		fmt.Printf("%s\n", output)
		input = readFromUser()
	}
	conn.Close()
}

func readFromUser() (output string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter message: ")
	text, _ := reader.ReadString('\n')

	fmt.Println(text)
	return text
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("Usage: ./client <ipandport>")
		os.Exit(1)
	}
	address := args[1]
	client(address)
}
