package main

import (
	"fmt"
	"golang.org/x/text/message"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	stats := getDiskSpaceStats()

	printer := message.NewPrinter(message.MatchLanguage("en"))
	printer.Printf("Total space: %.fMB\n", stats.TotalSpace)
	printer.Printf("Available space: %.fMB\n", stats.AvailableSpace)
}

type DiskSpaceStats struct {
	TotalSpace     float64
	AvailableSpace float64
}

func getDiskSpaceStats() DiskSpaceStats {
	output := getDiskSpaceOutput()

	lines := strings.Split(output, "\n")
	thirdLastLine := lines[len(lines)-3]
	columns := strings.Fields(thirdLastLine)

	return DiskSpaceStats{
		TotalSpace:     extractSizeInMb(columns[1]),
		AvailableSpace: extractSizeInMb(columns[3]),
	}
}

func getDiskSpaceOutput() string {
	cmd := exec.Command("df", "-h")
	stdout, err := cmd.Output()

	if err != nil {
		panic(err.Error())
	}

	return string(stdout)
}

func extractSizeInMb(sizeString string) float64 {
	sizeUnit, err := extractUnit(sizeString)

	if err != nil {
		panic(err)
	}

	// Invalid sizeUnit so return 0
	if sizeUnit == "" {
		return 0
	}

	nonDigitsRegex := regexp.MustCompile(`\D`)
	digits := nonDigitsRegex.ReplaceAllString(sizeString, "")

	// No digits so return 0
	if len(digits) == 0 {
		return 0
	}

	size, err := strconv.Atoi(digits)

	if err != nil {
		panic(err)
	}

	return sizeStringToMb(size, sizeUnit)
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
