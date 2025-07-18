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
		totalSpace := extractSizeInMb(columns[1])
		usedSpace := extractSizeInMb(columns[2])
		availableSpace := extractSizeInMb(columns[3])

		fmt.Println(fmt.Sprintf("totalSpace: %.fMB, usedSpace: %.fMB, availableSpace: %.fMB", totalSpace, usedSpace, availableSpace))
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

// Convert a string in format like 10Gi, 10Mi, or 10Ki to bytes
func extractSizeInMb(val string) float64 {
	nonDigitsRegex := regexp.MustCompile(`\D`)

	unit, err := extractUnit(val)

	if err != nil {
		panic(err)
	}

	// Invalid unit so return 0
	if unit == "" {
		return 0
	}

	digits := nonDigitsRegex.ReplaceAllString(val, "")

	// No digits so return 0
	if len(digits) == 0 {
		return 0
	}

	size, err := strconv.Atoi(digits)

	if err != nil {
		panic(err)
	}

	return sizeStringToMb(size, unit)
}

type SizeUnit string

const (
	Gi SizeUnit = "Gi"
	Mi SizeUnit = "Mi"
	Ki SizeUnit = "Ki"
)

func extractUnit(sizeString string) (unit SizeUnit, err error) {
	lastTwoChars := sizeString[len(sizeString)-2:]

	isValidUnit, err := regexp.MatchString("Gi|Mi|Ki", lastTwoChars)

	if err != nil {
		return "", err
	}

	if isValidUnit {
		return SizeUnit(lastTwoChars), nil
	}

	return "", nil

}

func sizeStringToMb(size int, unit SizeUnit) float64 {
	bytes := sizeStringToBytes(size, unit)

	return bytesToMb(bytes)
}

func sizeStringToBytes(size int, unit SizeUnit) int {
	switch unit {
	case Gi:
		return size * 1024 * 1024 * 1024
	case Mi:
		return size * 1024 * 1024
	case Ki:
		return size * 1024
	default:
		panic(fmt.Sprintf("Invalid unit: %s", unit))
	}
}

func bytesToMb(bytes int) float64 {
	return float64(bytes) / (1024 * 1024)
}
