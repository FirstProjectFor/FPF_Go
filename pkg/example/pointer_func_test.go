package demo

import (
	"fmt"
	"testing"
)

type P struct {
	method string
}

// p is a copy, not change field method's value
func (p P) changeValue() {
	p.method = "changeValue"
}

//p is a address, change field method's value
func (p *P) changeValuePointer() {
	p.method = "changeValuePointer"
}

func TestPointerFunc(t *testing.T) {
	p := P{method: "Empty"}
	fmt.Println("before changeValue:", p)
	p.changeValue()
	fmt.Println("after changeValue:", p)
	p.changeValuePointer()
	fmt.Println("after changeValue:", p)
}
