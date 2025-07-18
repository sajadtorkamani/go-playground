package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("df", "-h")

	stdout, err := cmd.Output()

	if err != nil {
		panic(err.Error())
	}

	cmdOutput := string(stdout)

	lines := strings.Split(cmdOutput, "\n")

	for index, line := range lines {
		fmt.Printf("Line %d: %s\n", index+1, line)
	}
}
