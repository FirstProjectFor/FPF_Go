package demo

import (
	"fmt"
	"os"
	"testing"
)

func TestPanic(t *testing.T) {

	defer func() {
		if r := recover(); nil != r {
			fmt.Println("recover")
		}
	}()

	_, err := os.Create("/tmp/file")

	if err != nil {
		panic(err)
	}

	panic("test")
}
