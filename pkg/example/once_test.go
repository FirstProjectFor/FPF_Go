package demo

import (
	"fmt"
	"sync"
	"testing"
)

var a string
var once sync.Once

func setup() {
	fmt.Println("set up execute")
	a = "hello, world"
}

func doprint() {
	once.Do(setup)
	fmt.Println(a)
}

func TestOnce(t *testing.T) {
	go doprint()
	go doprint()
}
