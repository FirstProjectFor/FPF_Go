package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {

	text := "Hello,世界"

	fmt.Println(text)
	fmt.Println(len(text))
	fmt.Println(utf8.RuneCountInString(text))

	for i := 0; i < len(text); {
		r, size := utf8.DecodeRuneInString(text[i:])
		fmt.Printf("%d,\t%c\n", i, r)
		i += size
	}

	for i, c := range text {
		fmt.Println("index: ", i, "char: ", string(c))
	}

	fmt.Println(string(30000))
	fmt.Println(string("30000"))
	fmt.Println(string(99999999999999999999999999))

	fmt.Printf("% x\n", text)

	r := []rune(text)
	fmt.Printf("% b\n", r)
	fmt.Printf("% d\n", r)
	fmt.Printf("% x\n", r)

	fmt.Println(unicode.IsLetter(9))
	fmt.Println(unicode.IsLetter('a'))
}
