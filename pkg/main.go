package main

import (
	"errors"
	"fmt"
)

func main() {

	err := errors.New("hello")

	fmt.Println(string([]byte(err.Error())))

}

type EE struct {

}

func (receiver *EE) Error() string {
	return ""
}