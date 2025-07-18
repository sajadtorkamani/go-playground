package jokes

import (
	"go-playground/api/db"
	"go-playground/api/models"
)

func JokeDoesNotExist(jokeContent string) bool {
	var joke models.Joke

	db.DB.Where("content = ?", jokeContent).First(&joke)

	return joke.ID == 0
}

func GetJokes() []models.Joke {
	var jokes []models.Joke
	db.DB.Find(&jokes)

	return jokes
}

func GetJokeCount() int64 {
	var totalNumJokes int64

	result := db.DB.Model(&models.Joke{}).Count(&totalNumJokes)

	if result.Error != nil {
		panic(result.Error)
	}

	return totalNumJokes
}
