package inner

import (
	"fmt"
	"testing"
)

type Name1 struct {
	name string
}

type Name2 struct {
	name string
}

type Person struct {
	Name1
	Name2
}

func (p Person) PrintName() {
	fmt.Println(p.Name1.name)
	fmt.Println(p.Name2.name)
}

func TestName(t *testing.T) {
	person := Person{}
	person.PrintName()
}
