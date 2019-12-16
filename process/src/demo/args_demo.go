package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	fmt.Println(os.Args)

	//从命令行接收参数
	//-name test -state 0 -age 37
	name := flag.String("name", "xiaotian", "user name")
	state := flag.Bool("state", true, "user state")
	age := flag.Int("age", 26, "user age")

	flag.Parse()

	fmt.Println("name: ", *name)
	fmt.Println("state: ", *state)
	fmt.Println("age: ", *age)

}
