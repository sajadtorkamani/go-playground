package main

import "fmt"

func main() {
	a, b := 10, 5

	if a > b {
		fmt.Printf("%v is greater than %v\n", a, b)
	} else {
		fmt.Printf("%v is less than %v\n", a, b)
	}
}
