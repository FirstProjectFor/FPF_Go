package demo

import (
	"regexp"
	"testing"
)

import "fmt"

import "bytes"

func TestRegex(t *testing.T) {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println("match: ", match)

	r := regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("peach: ", r.MatchString("peach"))
	fmt.Println("peaach: ", r.MatchString("peaach"))
	fmt.Println("peaachch  peaadasdaadasdasdh: ", r.FindString("peaachch  peaadasdch paadaschdasdh"))

	in := []byte("a peach a peachchchchhchch")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
