package main

import (
	"crypto/rand"
	"crypto/tls"
	"io"
	"log"
	"net"
	"time"
)

var (
	certFile = "public_key.crt"
	keyFile  = "private_key.pem"
)

func main() {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatalf("server:loadkeys:%s", err)
	}
	config := tls.Config{Certificates: []tls.Certificate{cert}}
	config.Time = time.Now
	config.Rand = rand.Reader
	service := "127.0.0.1:8000"
	listener, err := tls.Listen("tcp", service, &config)
	if err != nil {
		log.Fatal("server:listen:%s", err)
	}
	log.Print("server:listening....%s", service)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print("server:accept :%s", err)
			break
		}
		log.Printf("server:accepted from %s", conn.RemoteAddr())
		go handleClient(conn)
	}

}
func handleClient(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 512)
	for {
		log.Print("server: conn:waiting")
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Printf("server:conn :read :%s", err)
			}
			break
		}
		log.Printf("server:conn:echo %q\n", string(buf[:n]))
		n, err = conn.Write(buf[:n])
		log.Printf("server:conn:write %d bytes", n)
		if err != nil {
			log.Printf("server:conn:write:%s", err)
			break
		}
		log.Println("server:conn:closeted!")
	}
}
