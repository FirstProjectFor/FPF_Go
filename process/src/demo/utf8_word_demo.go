package main

import "bufio"

import "os"

import "io"

import "unicode"

import "unicode/utf8"

import "fmt"

func main() {

	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int

	fmt.Println(utflen)

	input := bufio.NewReader(os.Stdin)
	invalid := 0

	for {
		r, n, err := input.ReadRune()
		fmt.Println(r)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}

	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, l := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, l)
		}
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid utf-8 characters\n", invalid)
	}

}
