package demo

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestInterfaceType(t *testing.T) {

	var w io.Writer
	fmt.Printf("%T\n", w)
	w = os.Stdout
	fmt.Printf("%T\n", w)
	w = new(bytes.Buffer)
	fmt.Printf("%T\n", w)
	w = nil
	fmt.Printf("%T\n", w)
}
