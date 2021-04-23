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
	udpConnect, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8686,
	})
	if err != nil {
		log.Fatal(err)
	}
	_, err = udpConnect.Write([]byte("hello i am client"))
	if err != nil {
		log.Fatal(err)
	}

	bytes := make([]byte, 100)
	read, _ := udpConnect.Read(bytes)

	ipAndPort := string(bytes[0:read])
	log.Println("ipAndPort: ", ipAndPort)
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
