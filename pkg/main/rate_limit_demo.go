package main

import (
	"fmt"
	"time"
)

func main() {

	requests := make(chan int, 3)
	for i := 0; i < 3; i++ {
		requests <- i
	}
	close(requests)
	tict := time.Tick(time.Millisecond * 500)

	for r := range requests {
		<-tict
		fmt.Println("process request: ", r)
	}

	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond * 500) {
			burstyLimiter <- t
		}
	}()

	burstyRequest := make(chan int, 9)
	for i := 0; i < 9; i++ {
		burstyRequest <- i
	}
	close(burstyRequest)

	for req := range burstyRequest {
		<-burstyLimiter
		fmt.Println("requesst: ", req)
	}

}
