package main

import "fmt"

func main() {

	outF := outFunction()

	for i := 0; i < 10; i++ {
		fmt.Println(outF())
	}

}

func outFunction() func() int {
	number := 0
	return func() int {
		number++
		return number
	}
}
