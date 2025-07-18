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
	db := connectToDb()

	hokes := []string{"Joke 1", "Joke 2", "Joke 3"}
	numJokesAdded := 0

	for _, jokeContent := range hokes {
		addJoke(db, jokeContent)
		numJokesAdded += 1
	}

	if numJokesAdded > 0 {
		fmt.Println(
			fmt.Sprintf("✅  Added new %d hokes. We now have %d hokes in total.", numJokesAdded, getJokeCount(db)),
		)
	} else {
		fmt.Println("No hokes added.")
	}
}

func connectToDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("jokes.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to DB")
	}

	err = db.AutoMigrate(&Joke{})

	if err != nil {
		panic(err)
	}

	return db
}

func addJoke(db *gorm.DB, jokeContent string) {
	var joke Joke
	db.Where("content = ?", jokeContent).First(&joke)

	// If the joke has been added, skip it
	if joke.ID > 0 {
		return
	}

	db.Create(&Joke{Content: jokeContent})
}

func getJokeCount(db *gorm.DB) int64 {
	var totalNumJokes int64

	result := db.Model(&Joke{}).Count(&totalNumJokes)

	if result.Error != nil {
		panic(result.Error)
	}

	return totalNumJokes
}
