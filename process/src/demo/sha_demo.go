package main

import "crypto/sha1"

import "fmt"

func main() {

	h := sha1.New()

	hello := "Hello Go!"

	h.Write([]byte(hello))

	fmt.Println(hello)
	fmt.Printf("%x\n", h.Sum(nil))
}
