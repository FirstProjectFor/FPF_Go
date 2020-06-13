package main

import "fmt"

func main() {

	naturals := make(chan int64)
	squares := make(chan int64)

	var x int64 = 0

	go func() {
		for ; ; x++ {
			naturals <- x
		}
	}()

	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	for {
		fmt.Println(<-squares)
	}
}
