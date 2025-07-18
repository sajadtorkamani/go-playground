package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Joke struct {
	gorm.Model
	Content string
}

func main() {
	db, err := gorm.Open(sqlite.Open("jokes.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to DB")
	}

	err = db.AutoMigrate(&Joke{})

	if err != nil {
		panic(err)
	}

	jokes := []string{"Joke 1", "Joke 2", "Joke 3"}

	for _, joke := range jokes {
		db.Create(&Joke{Content: joke})
	}

	fmt.Println(fmt.Sprintf("Added %d jokes", len(jokes)))
}
