package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Duration(3) * time.Second)
		c1 <- "c1"
	}()

	go func() {
		time.Sleep(time.Duration(1) * time.Second)
		c2 <- "c2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("msg1: " + msg1)
		case msg2 := <-c2:
			fmt.Println("msg2: " + msg2)
		}
	}
}
