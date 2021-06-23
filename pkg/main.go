package main

import (
	"fmt"
	"time"
)

func main() {

	closeChain := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		close(closeChain)
	}()

	for {
		select {
		case <-closeChain:
		default:
			fmt.Println("default")
			continue
		}
		fmt.Println("close")
	}
}
