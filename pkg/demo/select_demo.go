package main

import (
	"fmt"
)

func main() {

	signal := make(chan bool, 1)

	go func() {
		for {
			signal <- true
		}
	}()

	for {
		select {
		case sig := <-signal:
			fmt.Println(sig)
		default:
			fmt.Print("default")
		}
	}
}
