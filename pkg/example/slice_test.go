package demo

import (
	"fmt"
	"strings"
	"testing"
)

func TestEmpty(t *testing.T) {
	var names []string
	names = append(names, "GO")
	fmt.Println(names)
}

func TestAppend(t *testing.T) {
	s := make([]string, 5)
	//empty s :  ,,,,
	fmt.Println("empty s : ", strings.Join(s, ","))
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	//s :  a,b,c,,
	fmt.Println("s : ", strings.Join(s, ","))
	//length:  5
	fmt.Println("length: ", len(s))

	s = append(s, "d")
	s = append(s, "e")
	s = append(s, "f")
	//s :  a,b,c,,,d,e,f
	fmt.Println("s : ", strings.Join(s, ","))
	//length:  8
	fmt.Println("length: ", len(s))
	//s[2,3] :  c
	fmt.Println("s[2,3] : ", strings.Join(s[2:3], ","))
	//s[2,6] :  c,,,d
	fmt.Println("s[2,6] : ", strings.Join(s[2:6], ","))
}

// why? see https://yourbasic.org/golang/gotcha-append/
func TestAppendCacheProblem(t *testing.T) {

	//b
	a := []byte("b")
	fmt.Printf("val: %s, len：%d, cap: %d \n", string(a), len(a), cap(a))

	aa := append(a, 'a')
	//ba
	fmt.Printf("val: %s, len：%d, cap: %d \n", string(aa), len(aa), cap(aa))

	b := append(aa, 'd')
	c := append(aa, 'g')

	//bag
	fmt.Printf("val: %s, len：%d, cap: %d \n", string(b), len(b), cap(b))
	//bag
	fmt.Printf("val: %s, len：%d, cap: %d \n", string(c), len(c), cap(c))
}