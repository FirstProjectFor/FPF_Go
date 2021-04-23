package demo

import (
	"fmt"
	"log"
	"net"
	"testing"
	"time"
)

func TestHttpServer(t *testing.T) {
	udpConnect, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8686,
	})
	if err != nil {
		log.Fatal(err)
	}
	for {
		bytes := make([]byte, 100)
		_, remoteAddr, err := udpConnect.ReadFromUDP(bytes)
		fmt.Println("revive data: ", string(bytes))
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("remote addr: ", remoteAddr.String())
		_, err = udpConnect.WriteToUDP([]byte(remoteAddr.String()), remoteAddr)
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(time.Second)
		//
		go func(remote *net.UDPAddr) {
			connect, err := net.DialUDP("udp4", nil, remote)
			if err != nil {
				log.Fatal(err)
			}
			_, err = connect.Write([]byte("hello i am server"))
			if err != nil {
				log.Fatal(err)
			}
		}(remoteAddr)
	}
}
