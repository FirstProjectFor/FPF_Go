package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker id: ", id, "processing job: ", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for j := 0; j < 3; j++ {
		go worker(j, jobs, results)
	}

	for i := 0; i < 9; i++ {
		jobs <- i
	}
	close(jobs)

	for k := 0; k < 9; k++ {
		<-results
	}
	close(results)

}
