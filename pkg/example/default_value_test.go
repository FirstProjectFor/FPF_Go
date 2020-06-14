package demo

import (
	"fmt"
	"testing"
)

func TestDefaultValue(t *testing.T) {
	fmt.Println("go" + "lang")
	fmt.Println("1 + 1 = ", 1+1)
	fmt.Println("7.0 / 3.0 = ", 7.0/3.0)
	fmt.Println("true && false", true && false)
	fmt.Println("true || false", true || false)
	fmt.Println("!true", !true)
}
