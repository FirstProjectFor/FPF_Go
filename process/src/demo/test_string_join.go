package main

import "time"

import "fmt"

func main() {

	count := 100000

	tempStr := ""
	value := "a"
	start := time.Now().UnixNano()
	for count > 0 {
		count--
		tempStr += value
	}
	end := time.Now().UnixNano()

	fmt.Println(end - start)
}
