package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
)

const (
	readWidth = 1024
	sep       = '\n'
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment the code below to pass the first stage
	//
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	conn, err := l.Accept()

	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	defer conn.Close()
	data := make([]byte, readWidth)
	bread, err := conn.Read(data)

	if err != nil {
		fmt.Println("Error reading from connection: ", err.Error())
		os.Exit(1)
	}
	if bread > 0 {
		ind := 0
		for ind > -1 {
			ind = bytes.IndexByte(data, sep)
			fmt.Println("message from connection: ", string(data[:ind]))
			if ind > -1 {
				conn.Write([]byte("+PONG\r\n"))
				data = data[ind+1:]
			}

		}

	}

}
