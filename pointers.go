package main

import "fmt"

func main() {
	var myPointer *int
	fmt.Println(myPointer)

	num := 5

	myPointer = &num

	fmt.Println(num)
	fmt.Println(&num)

	fmt.Println(myPointer)
	fmt.Println(*myPointer)
}
