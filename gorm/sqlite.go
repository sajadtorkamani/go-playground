package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Joke struct {
	gorm.Model
	content string
}

func main() {
	db, err := gorm.Open(sqlite.Open("jokes.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to DB")
	}

	db.AutoMigrate(&Joke{})

	db.Create(&Joke{content: "Some funny hoke..."})
	db.Create(&Joke{content: "Another funny joke..."})

	fmt.Println("Connected to db successfully")
}
