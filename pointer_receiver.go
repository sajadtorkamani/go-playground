package main

import "fmt"

type Book struct {
	title string
}

func (book Book) ChangeNameV1(newTitle string) {
	book.title = newTitle
}

func (book *Book) ChangeNameV2(newTitle string) {
	book.title = newTitle
}

func main() {
	book1 := Book{title: "War and Hate"}
	book1.ChangeNameV1("Peace and Love")
	// Outputs 'War and Hate' because the receiver is a value (not a pointer)
	// and so the `ChangeNameV1` method operates on a copy of `book1`, not the
	//`book1` itself.
	fmt.Println(book1)

	book2 := Book{title: "Burger and Coke"}
	book2.ChangeNameV2("Salad and Water")

	// Outputs 'Salad and Water' because receiver is a pointer and so the
	// `ChangeNameV2` method operates on the actual value of `book2`
	fmt.Println(book2)
}
