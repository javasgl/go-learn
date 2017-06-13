package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	conn, err := net.Dial("tcp", service)
	defer conn.Close()
	checkError(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.1\r\n\r\n"))
	checkError(err)
	result, err := readFully(conn)
	result2, err := ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println(string(result))
	fmt.Println(string(result2))
	os.Exit(0)
}

func readFully(conn net.Conn) ([]byte, error) {
	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s ", err.Error())
		os.Exit(1)
	}
}
