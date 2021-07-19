package main

import (
	"fmt"
)

type Animal struct {
	Name string
	Age  int
}

func (p *Animal) GetName() string {
	return p.Name
}

func (p *Animal) GetAge() int {
	return p.Age
}

type zoo struct {
	*Animal
}

func (z *zoo) GetName() string {
	return "zoo"
}

func (z *zoo) GetAge() int {
	return 0
}

func main() {
	animal := &Animal{
		Age:  20,
		Name: "animal",
	}
	printAnimal(animal)

	//z := zoo{
	//	animal,
	//}
	////Cannot use 'z' (type zoo) as type *Animal
	//printAnimal(z)
}

func printAnimal(animal *Animal) {
	fmt.Println(animal.GetName(), animal.GetAge())
}
