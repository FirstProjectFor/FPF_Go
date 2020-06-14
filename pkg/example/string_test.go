package demo

import (
	"fmt"
	"testing"
)

func TestStringArr(t *testing.T) {

	name := "张三"
	charArr := []rune(name)
	charArr[0] = '李'
	charArr[1] = '四'

	//李四
	fmt.Println(string(charArr))

	byteArr := []byte(name)
	//err  constant 26446 overflows byte
	//byteArr[0] = '李'
	fmt.Println(string(byteArr))
}
