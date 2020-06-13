package main

import (
	"fmt"
)

type fruit struct {
	name string
	c    color
}

type color struct {
	name string
}

func (f fruit) String() string {
	return fmt.Sprint("fruit:", f.name, " color: ", f.c)
}

func (f fruit) changeNameUseReference() {
	f.name = "f_c_r"
}

func (f *fruit) changeNameUsePoint() {
	f.name = "f_c_p"
}

func (c color) String() string {
	return c.name
}

//参数传递为引用传递,方法内部接收到的参数是原有变量的一个复制,对应有一个新的内存地址,
//在方法内部修改变量属性之后，原来方法外部的数据内容不受影响
//当传递的参数为为指针的时候,实际传递的是数据的地址,
//因此在方法内部对参数的修改会影响到原有变量的内容。
func main() {
	//测试参数为指针和引用
	f := fruit{}
	f.name = "f"
	f.c.name = "f"
	//fruit:f color: f
	fmt.Println("Before Change: ", f)

	changeUseReference(f)
	//fruit:f color: f
	fmt.Println("After Change User Reference: ", f)

	changeUsePoint(&f)
	//fruit:f_p color: f_c_p
	fmt.Println("After Change User Pointer: ", f)
	//初始化
	f = fruit{}
	f.name = "f"
	f.c.name = "f"
	//fruit:f color: f
	fmt.Println("Before Change: ", f)
	f.changeNameUseReference()
	//fruit:f color: f
	fmt.Println("After Change User Reference: ", f)
	f.changeNameUsePoint()
	//fruit:f_c_p color: f
	fmt.Println("After Change User Pointer: ", f)

	//测试参数为slice
	var names []string
	afterAppend := append(names, "zs")
	fmt.Println(names)       //[]
	fmt.Println(afterAppend) // [zs]

	//测试map
	labelCodeMap := map[string]int{}
	labelCodeMap["a"] = 1
	//map[a:1]
	fmt.Println(labelCodeMap)
	addLabel(labelCodeMap, "b", 2)
	//map[a:1 b:2]
	fmt.Println(labelCodeMap)
}

func changeUseReference(f fruit) {
	f.name = "f_r"
	f.c.name = "f_c_r"
}

func changeUsePoint(f *fruit) {
	f.name = "f_p"
	f.c.name = "f_c_p"
}

func addLabel(m map[string]int, key string, code int) {
	m[key] = code
}
