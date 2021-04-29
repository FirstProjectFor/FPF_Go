package main

import "fmt"

func main() {
	p := &Person{
		name: "1",
	}

	fmt.Println(p)
	changeName(p)
	fmt.Println(p)
}

func changeName(person *Person) {
	person.name = "2"
}

type Person struct {
	name string
}
