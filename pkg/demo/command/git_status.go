package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	command := exec.Command("git", "status")
	command.Stdout = os.Stdout
	if err := command.Run(); nil != err {
		fmt.Println(err)
	}
	fmt.Println(command.ProcessState.ExitCode())
}
