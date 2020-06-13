package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			defer f.Close()
			if nil != err {
				fmt.Fprintf(os.Stderr, "dub2: %v\n", err)
				continue
			}
			countLines(f, counts)
		}
	}

	for line, count := range counts {
		fmt.Println(line, ":", count)
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		text := input.Text()
		counts[text]++
	}
}
