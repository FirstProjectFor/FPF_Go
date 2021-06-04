package main

import (
	"fmt"
	"time"
)

func main() {
	defer p()
	time.Sleep(time.Second * 10)
}

func p() {
	fmt.Println("a")
}
