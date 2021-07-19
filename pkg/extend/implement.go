package main

import (
	"fmt"
	"reflect"
)

type Fruit interface {
	Name() string
}

type Apple struct {
}

func (a *Apple) Name() string {
	return "Apple"
}

type WaterMelon struct {
}

func (a *WaterMelon) Name() string {
	return "WaterMelon"
}

func showName(f Fruit) {
	fType := reflect.TypeOf(f)
	fmt.Println(fType)
	fmt.Println(f.Name())
	fmt.Println()
}

func main() {
	apple := Apple{}
	waterMelon := WaterMelon{}

	showName(&apple)
	showName(&waterMelon)
}
