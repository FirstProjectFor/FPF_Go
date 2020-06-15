package demo

import (
	"os"
	"testing"
)

import "fmt"

func TestDefer(t *testing.T) {
	file := createfile("/tmp/test.txt")
	defer closeFile(file)
	writeFiel(file)
}

func createfile(p string) *os.File {
	fmt.Println("create file: ", p)
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFiel(file *os.File) {
	fmt.Println("write to file: ", file)
	fmt.Fprintln(file, "data")
}

func closeFile(file *os.File) {
	fmt.Println("close file: ", file)
	file.Close()
}
