package demo

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

//从命令行接收参数
//-name test -state 0 -age 37
func TestFlag(t *testing.T) {

	fmt.Println(os.Args)

	name := flag.String("name", "xiaotian", "user name")
	state := flag.Bool("state", true, "user state")
	age := flag.Int("age", 26, "user age")

	flag.Parse()
	fmt.Println("name: ", *name)
	fmt.Println("state: ", *state)
	fmt.Println("age: ", *age)
}
