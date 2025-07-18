package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	currentDirectory, err := os.Getwd()

	if err != nil {
		fmt.Println("Failed to determine current directory")
		return
	}

	file, err := os.ReadFile(fmt.Sprintf("%s/data/quotes.txt", currentDirectory))

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(file), "\n")

	for i := 0; i < len(lines); i++ {
		fmt.Println(lines[i])
	}
}
