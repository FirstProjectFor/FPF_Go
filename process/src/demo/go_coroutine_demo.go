package main

import (
	"fmt"
	"time"
)

func f1(msg string) {
	for i := 0; i < 3; i++ {
		showMsg := fmt.Sprintf("msg: %s, index: %d", msg, i)
		fmt.Println(showMsg)
	}
}

func main() {
	f1("hello")

	go f1("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("inner func")

	time.Sleep(time.Duration(1) * time.Second)
	fmt.Println("done")
}
