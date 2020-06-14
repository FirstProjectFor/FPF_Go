package demo

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestScanner(t *testing.T) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		t := strings.ToUpper(scanner.Text())
		fmt.Println(t)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
