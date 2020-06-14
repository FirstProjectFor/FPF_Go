package demo

import (
	"fmt"
	"testing"
)

//InitNum func
func InitNum() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func TestClosePackage(t *testing.T) {

	numGenerator := InitNum()
	fmt.Println("num 0", numGenerator())
	fmt.Println("num 0", numGenerator())

	numGenerator1 := InitNum()
	fmt.Println("num 0", numGenerator1())
	fmt.Println("num 0", numGenerator1())

}
