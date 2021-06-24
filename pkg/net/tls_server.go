package main

import (
	"bufio"
	"crypto/tls"
	"github.com/xiaotian/demo/pkg/net/common"
	"log"
	"net"
)

type TLSServer struct {
}

func (t *TLSServer) start() {
	keyPair, err := tls.LoadX509KeyPair(common.CertFile, common.KeyFile)
	if err != nil {
		panic(err)
	}

	config := &tls.Config{Certificates: []tls.Certificate{keyPair}}

	ln, err := tls.Listen("tcp", common.HttpsListenAddr, config)
	if err != nil {
		log.Println(err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	for {
		msg, err := r.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}

		println(msg)

		n, err := conn.Write([]byte("world\n"))
		if err != nil {
			log.Println(n, err)
			return
		}
	}
}

func main() {
	server := TLSServer{}
	server.start()
}
