package main

import (
	"fmt"
	"net"
	"os"
	"strconv"

	resp "github.com/codecrafters-io/redis-starter-go/internal"
)

const (
	readWidth = 1024
	sep       = '\n'
)

var rn = []byte{'\r', '\n'}

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
	for {
		bread, _ := conn.Read(data)
		if bread == 0 {
			break
		}
		fmt.Println("data: ", string(data))
		data = data[:bread]

		arr := resp.Parse(&data)
		fmt.Println("the number of values in returned arr: ", len(arr))
		for _, v := range arr {
			if v == "PING" {
				sendPong(conn)
			} else {

				conn.Write([]byte{'$'})
				conn.Write([]byte(strconv.Itoa(len(v))))
				conn.Write(rn)
				conn.Write([]byte(v))
				conn.Write(rn)
			}

		}
	}

}

func sendPong(conn net.Conn) {
	conn.Write([]byte("+PONG\r\n"))
}
