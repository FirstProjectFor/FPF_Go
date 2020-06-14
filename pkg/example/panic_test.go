package demo

import (
	"os"
	"testing"
)

func TestPanic(t *testing.T) {

	_, err := os.Create("/tmp/file")

	if err != nil {
		panic(err)
	}

	panic("test")
}
