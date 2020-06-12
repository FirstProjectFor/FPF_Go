package main

import "fmt"

// why? see https://yourbasic.org/golang/gotcha-append/
func main() {

	//b
	a := []byte("b")
	fmt.Printf("val: %s, len：%d, cap: %d \n", string(a), len(a), cap(a))

	aa := append(a, 'a')
	//ba
	fmt.Printf("val: %s, len：%d, cap: %d \n", string(aa), len(aa), cap(aa))

	b := append(aa, 'd')
	c := append(aa, 'g')

	//bag
	fmt.Printf("val: %s, len：%d, cap: %d \n", string(b), len(b), cap(b))
	//bag
	fmt.Printf("val: %s, len：%d, cap: %d \n", string(c), len(c), cap(c))
}
