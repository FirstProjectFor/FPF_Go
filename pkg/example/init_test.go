package demo

import (
	"fmt"
	"testing"
)

func init() {
	fmt.Println("init 1")
}

func init() {
	fmt.Println("init 2")
}

func init() {
	fmt.Println("init 3")
}

func TestInit(t *testing.T) {

}
