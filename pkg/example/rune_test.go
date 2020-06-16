package demo

import (
	"fmt"
	"reflect"
	"testing"
	"unicode/utf8"
)

func TestStringTORune(t *testing.T) {

	content := "Hello,世界"

	for i := 0; i < len(content); {
		s, size := utf8.DecodeRuneInString(content[i:])
		fmt.Println("rune:", s, " char(s): ", string(s), " size: ", size, " type: ", reflect.TypeOf(s))
		i = i + size
	}

	for i, s := range content {
		fmt.Printf("%d, %q, %d \n", i, s, s)
	}
}
