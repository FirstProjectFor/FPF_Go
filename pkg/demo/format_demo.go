package main

import "fmt"

import "math"

type point struct {
	x int
	y int
}

func main() {

	p := point{x: 1, y: 2}

	fmt.Printf("p -> %s: %v\n", "%v", p)
	fmt.Printf("p -> %s: %+v\n", "%+v", p)
	fmt.Printf("p -> %s: %#v\n", "%#v", p)
	fmt.Printf("p -> %s: %T\n", "%T", p)

	fmt.Printf("true -> %s: %t\n", "%t", true)

	fmt.Printf("test -> %s: %s\n", "%s", "test")
	fmt.Printf("test -> %s: %x\n", "%x", "test")
	fmt.Printf("100 -> %s: %d\n", "%d", 100)
	fmt.Printf("100 -> %s: %b\n", "%b", 100)
	fmt.Printf("Pi -> %s: %f\n", "%f", math.Pi)
	fmt.Printf("Pi -> %s: |%6.0f|\n", "%6f", math.Pi)
	fmt.Printf("Pi -> %s: |%6.3f|\n", "%6.6f", math.Pi)
	fmt.Printf("Pi -> %s: |%-6.0f|\n", "%-6f", math.Pi)
	fmt.Printf("Pi -> %s: |%-6.3f|\n", "%-6.6f", math.Pi)

}
