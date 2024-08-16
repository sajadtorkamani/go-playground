package main

import "fmt"

func main() {
	names := [4]string{"John", "Bob", "Jill", "Tim"}
	firstTwo := names[0:2]

	fmt.Println(names)

	firstTwo[0] = "Mike"
	fmt.Println(firstTwo)
}
