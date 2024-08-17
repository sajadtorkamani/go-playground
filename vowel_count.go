package main

import (
	"fmt"
	"regexp"
)

func VowelCount(str string) int {
	var vowels []string
	regex := regexp.MustCompile(`[aeiou]`)
	matches := regex.FindAllStringSubmatch(str, -1)

	for _, char := range matches {
		vowels = append(vowels, char[0])
	}

	return len(vowels)
}

func main() {
	fmt.Println(VowelCount("sajad torkamani"))
}
