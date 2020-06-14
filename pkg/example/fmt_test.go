package demo

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

func TestFmt(t *testing.T) {

	sum1 := sha256.Sum256([]byte{'x'})
	sum2 := sha256.Sum256([]byte{'X'})
	result := fmt.Sprintf("sum1: %x\nsum2: %x\neq: %t\ntype: %T\n", sum1, sum2, sum1 == sum2, sum1)
	fmt.Println(result)

	bytes := [1]int{1}
	changeByte(bytes)
	fmt.Println(bytes)

	changeBytePtr(&bytes)
	fmt.Println(bytes)

	fmt.Printf("slice: %T\n", []int{})
	fmt.Println(cap([]int{1, 2, 3}))
}

func changeByte(bytes [1]int) {
	bytes[0] = -1
}

func changeBytePtr(ptr *[1]int) {
	ptr[0] = -2
}
