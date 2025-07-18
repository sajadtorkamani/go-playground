package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("whoami")

	stdout, err := cmd.Output()

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("You are %s", string(stdout))
}
