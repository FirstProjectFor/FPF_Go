package demo

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"testing"
)

func TestClient(t *testing.T) {
	conn, err := net.Dial("tcp", ":8686")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	fmt.Fprintf(conn, "hello\n")
	res, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(res))
	fmt.Fprintf(conn, "Jimmy!\n")
}
