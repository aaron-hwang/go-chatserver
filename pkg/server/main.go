package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
)

func server() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("Error creating tcp server")
		os.Exit(1)
	}
	fmt.Println("Listening on " + ln.Addr().String())

	// Accept and handle connections
	for {
		conn, ok := ln.Accept()
		if ok != nil {
			fmt.Printf("Error when trying to accept connection")
		}
		fmt.Println("Accepted connection!")
		handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	output := readFromClient(conn)
	for output != "quit" {
		fmt.Println("Read: " + output)
		output = readFromClient(conn)
	}
	conn.Close()
}

func readFromClient(conn net.Conn) (output string) {
	reader := bufio.NewReader(conn)
	out, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error when reading from client, %s", err)
		conn.Close()
		os.Exit(1)
	}
	conn.Write([]byte("Read a meessage of length: " + strconv.Itoa(len(out))))
	return string(out[:])
}

func main() {
	server()
}
