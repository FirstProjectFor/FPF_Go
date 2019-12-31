package main

import "fmt"

import "os"

func main() {

	fmt.Println("env demo")

	os.Setenv("name", "xiaotian")
	fmt.Println("env name= ", os.Getenv("name"))

	for n, v := range os.Environ() {
		fmt.Println("name: ", n, "value: ", v)
	}
}
