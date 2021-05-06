package main

import (
	"fmt"
	"github.com/hashicorp/mdns"
	"time"
)

func main() {
	// Make a channel for results and start listening
	entriesCh := make(chan *mdns.ServiceEntry, 4)
	go func() {
		for entry := range entriesCh {
			fmt.Printf("Got new entry: %v\n", entry)
		}
	}()

	// Start the lookup
	err := mdns.Lookup("_foobar._tcp", entriesCh)
	if err != nil {
		panic(err)
	}
	time.Sleep(time.Minute * 10)
}
