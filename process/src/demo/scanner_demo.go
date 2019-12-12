package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		t := strings.ToUpper(scanner.Text())
		fmt.Println(t)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
