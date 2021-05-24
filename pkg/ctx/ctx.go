package main

import (
	"fmt"
	"math/rand"
	"time"
)

const sleepTime = 10

func main() {
	for i := 0; i < 50; i++ {
		data := queryData()
		fmt.Println(data)
	}
}

type data struct {
}

func queryData() *data {
	resultChan := make(chan *data)
	timer := time.NewTimer(time.Millisecond * sleepTime)
	go query(resultChan)
	select {
	case result := <-resultChan:
		return result
	case <-timer.C:
		return nil
	}
}

func query(ch chan<- *data) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(sleepTime*2)))
	ch <- &data{}
}
