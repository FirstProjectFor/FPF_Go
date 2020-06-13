package main

import (
	"fmt"
	"sort"
	"sync"
)

// why? see https://yourbasic.org/golang/gotcha-data-race-closure/
func main() {
	var wg sync.WaitGroup
	times := 50
	numbers := make(chan int, times)
	wg.Add(times)
	for i := 0; i < times; i++ {
		go func(n int) {
			numbers <- n
			fmt.Print(" ", n)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println()
	close(numbers)

	result := make([]int, times)
	index := 0
	for n := range numbers {
		result[index] = n
		index++
	}
	sort.Ints(result)

	fmt.Println(result)
}
