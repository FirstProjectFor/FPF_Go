package main

import "fmt"

import "math"

func main() {

	number := 1000

	for i := 0; i < number; i++ {
		fmt.Printf("sin(%d) = %f\n", i, math.Sin(float64(i)))
	}

}
