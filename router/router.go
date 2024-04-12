package router

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"main.go/handlers"
)

func Run() {
	router := gin.New()

	// Implement the GET method
	router.GET("/persons", handlers.ListPersonsHandler)
	router.GET("/persons/:id", handlers.GetPersonDetails)

	// Implement the POST method
	router.POST("/persons", handlers.CreatePersonsHandler)

	// Implement the DELETE method
	router.DELETE("/persons/:id", handlers.DeletePersonsHandler)

	//Implement the PUT method
	router.PUT("/persons/:id", handlers.UpdatePersonHandler)

	port := os.Getenv("PORT")
	fmt.Println("Server listening on :" + port)
	if err := router.Run(":" + port); err != nil {
		fmt.Printf("Error on starting server: %v", err)
	}
}
