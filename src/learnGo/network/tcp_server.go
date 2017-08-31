package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	service := ":1200"

	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	CheckError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	CheckError(err)

	for {

		conn, err := listener.Accept()

		if err != nil {
			continue
		}

		go HandlerConn(conn)

	}

}

func CheckError(err error) {

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

}

func HandlerConn(conn net.Conn) {
	defer conn.Close()
	daytime := time.Now().String()
	conn.Write([]byte(daytime))
}
