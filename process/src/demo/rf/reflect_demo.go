package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {

	so := os.Stdout

	t := reflect.TypeOf(so)

	fmt.Println(t.String())
}
