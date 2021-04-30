package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("quotes.txt")

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(file), "\n")

	for i := 0; i < len(lines); i++ {
		fmt.Println(lines[i])
	}
}
