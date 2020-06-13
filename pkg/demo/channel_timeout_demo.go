package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "msg1"
	}()

	select {
	case msg1 := <-c1:
		fmt.Println("msg1: " + msg1)
	case <-time.After(time.Second * 1):
		fmt.Println("get msg1 time out!")
	}

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "msg2"
	}()

	select {
	case msg2 := <-c1:
		fmt.Println("msg2: " + msg2)
	case <-time.After(time.Second * 2):
		fmt.Println("get msg1 time out!")
	}

}
