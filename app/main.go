package main

import (
	"fmt"
	"net"
	"os"

	resp "github.com/codecrafters-io/redis-starter-go/internal"
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

	sendResponse(l)

}

func sendResponse(l net.Listener) {
	conn, err := l.Accept()

	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	defer conn.Close()
	data := make([]byte, readWidth)

	bread, err := conn.Read(data)
	fmt.Println("data: ", string(data))
	arr := resp.Parse(data[:bread])
	fmt.Println("the number of values in returned arr: ", len(arr))
	for _, v := range arr {
		if v == "PING" {
			sendPong(conn)
		} else {
			conn.Write([]byte(v))
		}

	}
}

func sendPong(conn net.Conn) {
	conn.Write([]byte("+PONG\r\n"))
}
