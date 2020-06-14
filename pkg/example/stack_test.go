package demo

import (
	"fmt"
	"testing"
)

func f111(n int) {
	fmt.Println(n)
	f111(n + 1)
}

func TestStack(t *testing.T) {
	f111(1)
}
