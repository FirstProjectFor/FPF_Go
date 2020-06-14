package demo

import (
	"fmt"
	"testing"
)

func TestFor(t *testing.T) {

	i := 0
	for i < 2 {
		fmt.Println("i = ", i)
		i++
	}

	for j := 0; j < 2; j++ {
		fmt.Println("j = ", j)
	}

	k := 0
	for {
		fmt.Println("k = ", k)
		k++
		if k >= 2 {
			break
		}
	}

	for l := 0; l < 4; l++ {
		if l%2 == 0 {
			continue
		}
		fmt.Println("l = ", l)
	}

}
