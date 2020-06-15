package demo

import (
	"fmt"
	"reflect"
	"testing"
)

type myInt int

func TestTypeConversion(t *testing.T) {
	var i int = 1
	var mi myInt = 1

	isMyInt(i)
	isMyInt(mi)
}

func isMyInt(v interface{}) {
	if _, ok := v.(myInt); ok {
		fmt.Println("type is `myInt`.    ", " v: ", v, " real type: ", reflect.TypeOf(v).Name())
		return
	}
	fmt.Println("type is not `myInt`.", " v: ", v, " real type: ", reflect.TypeOf(v).Name())
}
