package resp

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

var sep = []byte{'\r', '\n'}

var cache map[string]string

const ECHO = "echo"

func init() {
	cache = make(map[string]string)
}

func Parse(datap *[]byte) []string {
	data := *datap
	fmt.Println("begin parsing")
	fmt.Println("len of data: ", len(data))
	parsed := make([]string, 0)
	switch data[0] {
	case '*':
		ind := bytes.Index(data, sep)
		if ind == -1 {
			fmt.Println("Improper message format, no endline characters found")
		} else {
			//l := getInt(data[1:ind])
			data = data[ind+2:]
			parsed = append(parsed, Parse(&data)...)
		}
	case '$':

		parsed = append(parsed, parseNextString(&data))
	case 'S':
		if string(data[0:3]) == "SET" {
			parsed = append(parsed, parseSetCommand(&data))
		}
	case 'G':
		if string(data[0:3]) == "GET" {
			parsed = append(parsed, parseGetCommand(&data))
		}
	}

	return parsed
}

func getInt(data []byte) int {
	val, err := strconv.Atoi(string(data))
	if err != nil {
		fmt.Printf("Could not convert int to string: %v", err)
		return 0
	}
	return val
}

func parseNextString(datap *[]byte) string {
	data := *datap
	ind := bytes.Index(data, sep)

	if ind == -1 {
		fmt.Println("Improper message format, no endline characters found")
	}
	l := getInt(data[1:ind])
	s := string(data[ind+2 : ind+2+l])
	ind = ind + 4 + l
	data = data[ind:]
	fmt.Println("current string: ", s)
	//*datap = data
	if strings.ToLower(s) == ECHO {
		s = parseNextString(&data)
	}
	return buildRespString(s)

}

func parseSetCommand(datap *[]byte) string {
	data := *datap
	ind := bytes.Index(data, sep)

	if ind == -1 {
		fmt.Println("Improper message format, no endline characters found")
	}
	sArr := strings.Split(string(data[:ind]), " ")

	cache[sArr[1]] = sArr[2]
	data = data[ind+2:]

	return "+OK\r\n"
}

func parseGetCommand(datap *[]byte) string {
	data := *datap
	ind := bytes.Index(data, sep)

	if ind == -1 {
		fmt.Println("Improper message format, no endline characters found")
	}
	sArr := strings.Split(string(data[:ind]), " ")

	data = data[ind+2:]
	if v, ok := cache[sArr[1]]; ok {
		return buildRespString(v)
	}
	return "-1"
}

func buildRespString(s string) string {
	if s == "-1" {
		return "$-1\r\n"
	}
	return fmt.Sprintf("$%d\r\n%s\r\n", len(s), s)
}
