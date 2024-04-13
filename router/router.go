package router

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"main.go/controllers"
	"main.go/handlers"
)

func Run() {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/signup", controllers.Signup)

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
