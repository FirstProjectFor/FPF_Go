package demo

import (
	"fmt"
	"testing"
)

type person1 struct {
	name string
	age  int
}

func TestStruct(t *testing.T) {
	p1 := person1{name: "xiaotian", age: 26}
	fmt.Println(p1)
	fmt.Println(p1.name)
	fmt.Println(p1.age)

	p2 := &p1
	fmt.Println(p1.age)
	p2.age = 27
	fmt.Println(p2.age)
	fmt.Println(p1.age)
}
