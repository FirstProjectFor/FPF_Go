package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	bytes := []byte("A")
	fmt.Println(len(bytes))
	bytes = []byte("ADFZ")
	fmt.Println(len(bytes))

	s := "联通"
	bytes = []byte(s)
	fmt.Println(len(bytes))
	fmt.Println(utf8.RuneCountInString(s))
}
