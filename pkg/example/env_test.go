package demo

import (
	"fmt"
	"testing"
)

import "os"

func TestEnv(t *testing.T) {

	fmt.Println("env example")

	os.Setenv("name", "xiaotian")
	fmt.Println("env name= ", os.Getenv("name"))

	for n, v := range os.Environ() {
		fmt.Println("name: ", n, "value: ", v)
	}
}
