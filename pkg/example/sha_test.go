package demo

import (
	"crypto/sha1"
	"testing"
)

import "fmt"

func TestSha(t *testing.T) {

	h := sha1.New()

	hello := "Hello Go!"

	h.Write([]byte(hello))

	fmt.Println(hello)
	fmt.Printf("%x\n", h.Sum(nil))
}
