package demo

import (
	"fmt"
	"testing"
)

type rect struct {
	widht  int
	height int
}

// area
func (r *rect) area() int {
	return r.widht * r.height
}

//perim
func (r rect) perim() int {
	return 2 * (r.widht + r.height)
}

func TestFunc(t *testing.T) {

	r1 := rect{4, 5}
	fmt.Println("r1 area ", r1.area())
	fmt.Println("r1 perim ", r1.perim())

	r2 := &r1
	fmt.Println("r2 area ", r2.area())
	fmt.Println("r2 perim ", r2.perim())
}
