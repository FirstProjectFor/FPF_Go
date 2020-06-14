package demo

import (
	"fmt"
	"testing"
)

func TestInnerFunc(t *testing.T) {

	done := make(chan bool)

	outF := outFunction()

	for i := 0; i < 10; i++ {
		fmt.Println(outF())
	}

	go func() {
		for i := 0; i < 100000; i++ {
			fmt.Printf("go: %d\n", i)
		}
		done <- true
	}()

	<-done
}

func outFunction() func() int {
	number := 0
	return func() int {
		number++
		return number
	}
}
