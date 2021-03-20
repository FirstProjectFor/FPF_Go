package main

import "fmt"

func main() {
	b := make([]byte, 5, 6)

	b[2] = 123
	b = append(b, 2)
	b = append(b, 2)
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(cap(b))

}
