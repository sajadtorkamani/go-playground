package main

import "fmt"

func main() {
	//str := "hello 😊 sajad"
	str := "😊"

	for index := 0; index < len(str); index++ {
		fmt.Println(str[index])
	}
}
