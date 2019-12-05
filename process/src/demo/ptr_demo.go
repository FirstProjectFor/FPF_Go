package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {

	i := 1
	fmt.Println("i = ", i)
	zeroval(i)
	fmt.Println("after change i = ", i)
	fmt.Println("")

	j := 1
	fmt.Println("before change j ptr = ", &j)
	fmt.Println("before change j = ", i)
	zeroptr(&j)
	fmt.Println("after change j ptr= ", &j)
	fmt.Println("after change j = ", j)
}
