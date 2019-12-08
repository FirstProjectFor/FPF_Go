package main

import (
	"fmt"
	"time"
)

func main() {

	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("start job: ", j)
			} else {
				fmt.Println("all job has finished!")
				done <- true
				return
			}
		}
	}()

	for i := 0; i < 3; i++ {
		jobs <- i
		fmt.Println("send job i: ", i)
		time.Sleep(time.Second * 1)
	}

	close(jobs)

	fmt.Println("send all jobs")
	<-done
}
