package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("start work ...")
	time.Sleep(time.Duration(3) * time.Second)
	fmt.Println("finish work ...")
	done <- true
}

func main() {
	done := make(chan bool, 1)
	worker(done)
	<-done
}
