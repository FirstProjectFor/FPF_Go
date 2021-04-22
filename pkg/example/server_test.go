package demo

import (
	"fmt"
	"log"
	"net"
	"testing"
)

func TestHttpServer(t *testing.T) {
	l, err := net.Listen("tcp", ":8686")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Println("remote addr: ", conn.RemoteAddr())
	err := conn.Close()
	if err != nil {
		fmt.Println(err)
	}
}
