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

	db.Create(&Joke{Content: "Some funny hoke..."})
	db.Create(&Joke{Content: "Another funny joke..."})

	fmt.Println("Connected to DB successfully")
}
