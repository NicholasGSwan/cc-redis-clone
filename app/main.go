package main

import (
	"fmt"
	"net"
	"os"
	"strings"

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

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Could not accept new connection")
			os.Exit(1)
		}
		if conn == nil {
			break
		}
		go sendResponse(conn)
	}

}

func sendResponse(conn net.Conn) {
	defer conn.Close()

	for {
		data := make([]byte, readWidth)
		bread, _ := conn.Read(data)
		if bread == 0 {
			break
		}

		data = data[:bread]

		arr := resp.Parse(&data)

		for _, v := range arr {
			if v == "PING" {
				sendPong(conn)
			} else {

				valArr := strings.Split(v, "\r\n")
				fmt.Println("the number of values in returned valArr: ", len(valArr))

				for _, val := range valArr {
					fmt.Println("curr val: ", val)
					conn.Write([]byte(val))
					conn.Write(rn)
				}

			}

		}
	}

}

func sendPong(conn net.Conn) {
	conn.Write([]byte("+PONG\r\n"))
}
