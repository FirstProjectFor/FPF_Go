package demo

import (
	"fmt"
	"testing"
)

import "os"

func TestExit(t *testing.T) {
	defer fmt.Println("Exit Go!")
	os.Exit(3)
}
