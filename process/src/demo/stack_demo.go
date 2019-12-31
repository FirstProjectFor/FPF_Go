package main

import "fmt"

func f(n int) {
	fmt.Println(n)
	f(n + 1)
}

func main() {
	f(1)
}
