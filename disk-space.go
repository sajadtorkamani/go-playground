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

	//totalSpace := 0
	//availableSpace := 0

	cmdOutput := string(stdout)
	lines := strings.Split(cmdOutput, "\n")

	for _, line := range lines[1:] {
		line = strings.Trim(line, " ")

		if !isValidLine(line) {
			continue
		}

		columns := strings.Fields(line)
		size := columns[1]
		used := columns[2]
		available := columns[3]

		fmt.Println(fmt.Sprintf("size: %s, used: %s, available: %s", size, used, available))
	}
}

func isValidLine(line string) bool {
	if len(line) == 0 {
		return false
	}
	if strings.Contains(line, "Filesystem") {
		return false
	}

	return true
}

func getOrFallback[T any](slice []T, index int, fallback T) T {
	if index >= 0 && index < len(slice) {
		return slice[index]
	}

	return fallback
}
