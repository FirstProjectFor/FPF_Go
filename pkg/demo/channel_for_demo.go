package main

import "fmt"

func main() {
	queue := make(chan string, 2)

	queue <- "e1"
	queue <- "e2"
	close(queue)

	for element := range queue {
		fmt.Println("element: ", element)
	}
}
