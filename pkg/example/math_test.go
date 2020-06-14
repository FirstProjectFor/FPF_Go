package demo

import (
	"fmt"
	"testing"
)

import "math"

func TestSin(t *testing.T) {
	number := 1000
	for i := 0; i < number; i++ {
		fmt.Printf("sin(%d) = %f\n", i, math.Sin(float64(i)))
	}
}
