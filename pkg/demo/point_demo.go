package main

import "fmt"

func main() {

	var age int64
	age = 26

	var pointer *int64 = &age
	fmt.Println("address", pointer)
	fmt.Println("value", *pointer)
	*pointer = 33
	fmt.Println("value", *pointer)
}
