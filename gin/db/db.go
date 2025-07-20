package db

import (
	"go-playground/api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Joke = models.Joke

var DB *gorm.DB

func ConnectToDatabase() {
	db, err := gorm.Open(sqlite.Open("jokes.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to DB")
	}

	err = db.AutoMigrate(&models.Joke{})

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

func GetJokes() []Joke {
	var jokes []Joke
	DB.Find(&jokes)

	return jokes
}

func GetJokeCount() int64 {
	var totalNumJokes int64

	result := DB.Model(&models.Joke{}).Count(&totalNumJokes)

	if result.Error != nil {
		panic(result.Error)
	}

	return totalNumJokes
}

func GetJokeById(jokeId int) Joke {
	var jokeResult Joke
	DB.First(&jokeResult, jokeId)

	return jokeResult

}
