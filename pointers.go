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

	&myPointer = 100

	fmt.Println(num)

	var myInt = 42
	var myIntPointer = &myInt

	fmt.Println(*myIntPointer) // prints 42

	&myIntPointer = 21 // sets the value myInt to 42 through the myIntPointer

	fmt.Println(myInt) // prints 21
}
