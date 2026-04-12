package main

import (
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
	for {
		bread, err := conn.Read(data)

		if err != nil {
			fmt.Println("Error reading from connection: ", err.Error())
			os.Exit(1)
		}
		if bread > 0 {

			go sendPong(conn)

		}
	}

}

func sendPong(conn net.Conn) {
	conn.Write([]byte("+PONG\r\n"))
}
