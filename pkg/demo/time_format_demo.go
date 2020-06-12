package main

import "time"

import "fmt"

func main() {

	p := fmt.Println

	now := time.Now()
	p(now)
	p(now.Format(time.RFC1123))
}
