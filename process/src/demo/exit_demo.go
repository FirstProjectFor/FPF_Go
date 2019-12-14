package main

import "fmt"

import "os"

func main() {
	defer fmt.Println("Exit Go!")
	os.Exit(3)
}
