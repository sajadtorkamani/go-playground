package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (person Person) PrintDetails() string {
	return fmt.Sprintf("Name is %s, age is %d", person.name, person.age)
}

func main() {
	person1 := Person{name: "Bob", age: 20}
	person2 := Person{name: "Alice", age: 30}

	fmt.Println(person1.PrintDetails())
	fmt.Println(person2.PrintDetails())
}
