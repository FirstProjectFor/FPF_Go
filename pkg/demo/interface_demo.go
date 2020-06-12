package main

import "fmt"

type animal interface {
	getName() string
	getAge() int
}

type person struct {
	name string
	age  int
}

func (p person) getName() string {
	return p.name
}

func (p person) getAge() int {
	return p.age
}

func printAnimal(a animal) {
	fmt.Println("animal ", a)
	fmt.Println("animal's name", a.getName())
	fmt.Println("animal's age", a.getAge())
}

func main() {
	p1 := person{"xiaotian", 26}
	printAnimal(p1)
}
