package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Joke struct {
	gorm.Model
	Content string
}

var DB *gorm.DB

func ConnectToDatabase() {
	db, err := gorm.Open(sqlite.Open("jokes.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to DB")
	}

	err = db.AutoMigrate(&Joke{})

	if err != nil {
		panic(err)
	}

	DB = db
}

func JokeDoesNotExist(jokeContent string) bool {
	var joke Joke

	DB.Where("content = ?", jokeContent).First(&joke)

	return joke.ID == 0
}

func GetJokeCount() int64 {
	var totalNumJokes int64

	result := DB.Model(&Joke{}).Count(&totalNumJokes)

	if result.Error != nil {
		panic(result.Error)
	}

	return totalNumJokes
}
