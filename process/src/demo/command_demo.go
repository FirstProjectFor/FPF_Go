package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {

	cls2 := exec.Command("jps", "-v")
	clsO, err := cls2.Output()
	if nil != err {
		panic(err)
	}
	fmt.Println(string(clsO))

	path, _ := exec.LookPath("jps")
	fmt.Println(path)

	args := []string{"./jps.exe", "-v"}

	en := os.Environ()
	e := syscall.Exec(path, args, en)
	if nil != e {
		panic(e)
	}
}
