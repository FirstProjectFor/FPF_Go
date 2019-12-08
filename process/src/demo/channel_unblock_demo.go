package main

import (
	"fmt"
	"time"
)

func main() {

	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("msg: " + msg)
	default:
		fmt.Println("default")
	}

	msg := "msg2"
	select {
	case messages <- msg:
		fmt.Println("send message to messages")
	default:
		fmt.Println("un send message to messages")
	}

	select {
	case msg := <-messages:
		fmt.Println("accept message from messages: ", msg)
	case signal := <-signals:
		fmt.Println("accept signal from signales: ", signal)
	default:
		fmt.Println("not accept any message")
	}

	time.Sleep(time.Second * 2)

}
