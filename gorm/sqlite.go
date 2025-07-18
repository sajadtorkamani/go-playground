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

	var totalNumJokes int64
	result := db.Model(&Joke{}).Count(&totalNumJokes)

	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println(
		fmt.Sprintf("✅  Added new %d jokes. We now have %d jokes in total.", len(jokes), totalNumJokes),
	)
}
