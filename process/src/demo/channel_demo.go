package main

import "fmt"

func main() {

	message := make(chan string)

	go func() {
		message <- "ping"
		message <- "pong"
	}()

	fmt.Println("msg = ", <-message)
	fmt.Println("msg = ", <-message)
}
