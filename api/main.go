package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-playground/api/db"
	"go-playground/api/jokes"
	"go-playground/api/models"
)

func main() {
	db.ConnectToDatabase()

	JOKES := []string{"Joke 1", "Joke 2", "Joke 3", "Joke 4"}
	numJokesAdded := 0

	for _, joke := range JOKES {
		if jokes.JokeDoesNotExist(joke) {
			db.DB.Create(&models.Joke{Content: joke})
			numJokesAdded += 1
		}
	}

	if numJokesAdded > 0 {
		fmt.Println(
			fmt.Sprintf("✅  Added new %d jokes. We now have %d jokes in total.", numJokesAdded, jokes.GetJokeCount()),
		)
	} else {
		fmt.Println("No jokes added.")
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello!",
		})
	})

	router.GET("/jokes", func(c *gin.Context) {
		c.JSON(200, jokes.GetJokes())
	})

	fmt.Println("🚀 Server listening on http://localhost:8080")

	router.Run()
}
