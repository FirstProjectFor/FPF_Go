package main

import (
	"errors"
	"fmt"
)

func testError(arg int) (int, error) {
	if arg%2 == 0 {
		return -1, errors.New("测试错误")
	}
	
	return 1, nil
	
}

func main() {

	_, e1 := testError(2)
	if e1 != nil {
		fmt.Println("e1", e1)
	}

	_, e2 := testError(1)
	if e2 != nil {
		fmt.Println("e2", e2)
	}

	for i, v := range []int{10, 42} {
		fmt.Println(fmt.Sprintf("index: %d, value: %d", i, v))
	}
}
