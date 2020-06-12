package main

import "image/color"

import "fmt"

// Point x,y point
type Point struct {
	X float64
	Y float64
}

// ColorPoint cp
type ColorPoint struct {
	Point
	Color color.RGBA
}

// ColorPointP color point p
type ColorPointP struct {
	*Point
	Color color.RGBA
}

func main() {

	var cp ColorPoint
	cp.X = 10.0
	cp.Point.Y = 11.0

	fmt.Println(cp.X)
	fmt.Println(cp.Y)
	fmt.Printf("%v\n", cp)

	p1 := Point{X: 1.0, Y: 1.0}
	p2 := Point{X: 2.0, Y: 2.0}

	pp1 := ColorPointP{&p1, color.RGBA{255, 0, 0, 255}}
	pp2 := ColorPointP{&p2, color.RGBA{0, 0, 255, 255}}

	fmt.Println(pp1)
	fmt.Println(pp2)

	fmt.Println(*pp1.Point)

}
