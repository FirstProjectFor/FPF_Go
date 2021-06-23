package main

import "fmt"

type IAnimal interface {
	GetName() string
	GetAge() int
}

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

type Pig struct {
	*Animal
}

func (p *Pig) GetName() string {
	return "pig"
}

func (p *Pig) GetAge() int {
	return 30
}

type zoo struct {
	*Animal
	ian IAnimal
}

func main() {
	animal := &Animal{
		Age:  20,
		Name: "animal",
	}
	z := zoo{
		Animal: animal,
		ian:    animal,
	}
	fmt.Println(z.GetName(), z.GetAge())
	fmt.Println(z.ian.GetName(), z.ian.GetAge())

	pig := &Pig{
		animal,
	}
	fmt.Println(pig.GetName(), pig.GetAge())

	z = zoo{
		ian: pig,
	}
	fmt.Println(z.ian.GetName(), z.ian.GetAge())
}
