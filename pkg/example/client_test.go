package demo

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"testing"
)

func TestClient(t *testing.T) {
	udpAddr := &net.UDPAddr{
		IP:   net.IPv4(43, 225, 158, 197),
		Port: 8686,
	}
	log.Println("Dial up: ", udpAddr)
	udpConnect, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("write data")
	_, err = udpConnect.Write([]byte("hello i am client"))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("read data")
	bytes := make([]byte, 100)
	read, _ := udpConnect.Read(bytes)
	ipAndPort := string(bytes[0:read])
	log.Println("ipAndPort: ", ipAndPort)

	//listen
	listen(ipAndPort)
}

func listen(ipAndPort string) {
	split := strings.Split(ipAndPort, ":")

	ipStr := strings.Split(split[0], ".")
	ip1, _ := strconv.Atoi(ipStr[0])
	ip2, _ := strconv.Atoi(ipStr[1])
	ip3, _ := strconv.Atoi(ipStr[2])
	ip4, _ := strconv.Atoi(ipStr[3])

	port, _ := strconv.Atoi(split[1])

	udpConnect, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(byte(ip1), byte(ip2), byte(ip3), byte(ip4)),
		Port: port,
	})
	if err != nil {
		log.Fatal(err)
	}
	for {
		bytes := make([]byte, 100)
		_, remoteAddr, err := udpConnect.ReadFromUDP(bytes)
		if err != nil {
			fmt.Println(err)
			continue
		}
		log.Println("revive data: ", string(bytes), "remoteAddr:", remoteAddr)
	}
}
