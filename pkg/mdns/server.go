package main

import (
	"github.com/hashicorp/mdns"
	"os"
	"time"
)

func main() {
	// Setup our service export
	host, _ := os.Hostname()
	info := []string{"My awesome service"}

	service, _ := mdns.NewMDNSService(host, "_foobar._tcp", "", "", 8000, nil, info)

	// Create the mDNS server, defer shutdown
	mdns.NewServer(&mdns.Config{Zone: service})
	time.Sleep(time.Minute * 10)
}
