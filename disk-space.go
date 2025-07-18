package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
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
		totalSpace := extractSize(columns[1])
		usedSpace := extractSize(columns[2])
		availableSpace := extractSize(columns[3])

		fmt.Println(fmt.Sprintf("totalSpace: %d, usedSpace: %d, availableSpace: %d", totalSpace, usedSpace, availableSpace))
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

func extractSize(val string) int {
	nonDigitsRegex := regexp.MustCompile(`\D`)
	digits := nonDigitsRegex.ReplaceAllString(val, "")

	if len(digits) > 0 {
		size, err := strconv.Atoi(digits)

		if err != nil {
			panic(err)
		}

		return size
	}

	return 0
}

func getOrFallback[T any](slice []T, index int, fallback T) T {
	if index >= 0 && index < len(slice) {
		return slice[index]
	}

	return fallback
}
