package demo

import (
	"io"
	"net"
	"testing"
	"time"
)

func TestDateServer(t *testing.T) {
	listener, err := net.Listen("tcp", "10.0.8.33:8989")
	if nil != err {
		panic(err)
	}

	for {
		connect, err := listener.Accept()
		if nil != err {
			panic(err)
		}
		go handel(connect)
	}

}

func handel(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, time.Now().Format("2006-01-02 15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
