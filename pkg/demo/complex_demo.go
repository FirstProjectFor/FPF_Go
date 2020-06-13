package main

import "fmt"

import "math/cmplx"

func main() {
	var ca complex128 = complex(1, 2)
	var cb complex128 = complex(3, 4)

	var cc = ca * cb
	fmt.Println(cc)
	fmt.Println(real(cc))
	fmt.Println(imag(cc))

	x := 1 + 2i
	y := 3 + 4i

	z := x * y

	fmt.Println(z)
	fmt.Println(real(z))
	fmt.Println(imag(z))

	fmt.Println(cmplx.Sqrt(z))
}
