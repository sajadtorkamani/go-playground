package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-playground/api/api"
	"go-playground/api/db"
	"go-playground/api/models"
	"net/http"
)

func main() {
	db.ConnectToDatabase()

	seedDatabase()

	router := initialiseRouter()
	router.Run()
}

func initialiseRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", apiEntry)
	router.GET("/jokes", api.ListJokes)
	router.GET("/jokes/:jokeId", api.GetJoke)
	router.POST("/jokes", api.CreateJoke)

	fmt.Println("🚀 Server listening on http://localhost:8080")

	return router
}

func apiEntry(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello!"})
}

func seedDatabase() {
	JOKES := []string{"Joke 1", "Joke 2", "Joke 3", "Joke 4"}
	numJokesAdded := 0

	for _, joke := range JOKES {
		if db.JokeDoesNotExist(joke) {
			db.DB.Create(&models.Joke{Content: joke})
			numJokesAdded += 1
		}
	}

	if numJokesAdded > 0 {
		fmt.Println(
			fmt.Sprintf("✅  Added new %d jokes. We now have %d jokes in total.", numJokesAdded, db.GetJokeCount()),
		)
	} else {
		fmt.Println("No jokes added.")
	}
}
