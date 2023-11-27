package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello from main")
	router := gin.New()

	// Implement the GET method
	router.GET("/persons", listPersonsHandler)
	router.GET("/persons/:id", getPersonDetails)

	// Implement the POST method
	router.POST("/persons", createPersonsHandler)

	// Implement the DELETE method
	router.DELETE("/persons/:id", deletePersonsHandler)

	//Implement the PUT method
	router.PUT("/persons/:id", updatePersonHandler)

	router.Run(":3000")
}

func listPersonsHandler(ctx *gin.Context) {
	ctx.JSON(200, persons)
}

func getPersonDetails(ctx *gin.Context) {
	id := ctx.Param(("id"))

	for _, val := range persons {
		if val.ID == id {
			ctx.JSON(200, val)
			return
		}
	}
	
	ctx.JSON(http.StatusNotFound, gin.H{"message": "Person not found"})
}

func createPersonsHandler(ctx *gin.Context){
	var person Person

	if err := ctx.ShouldBindJSON(&person); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	persons = append(persons, person)
	ctx.JSON(http.StatusCreated, gin.H{"message": "Person details created"})
}

func deletePersonsHandler(ctx *gin.Context){
	id := ctx.Param("id")

	for i, a := range persons {
		if a.ID == id {
			persons = append(persons[:i], persons[i+1:]...)
			ctx.JSON(200, gin.H{"message": "Person removed"})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"Message": "Person not found"}) // 204 No Content
}

func updatePersonHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	var updatePerson Person

	if err := ctx.ShouldBindJSON(&updatePerson); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, person := range persons {
		if person.ID == id {
			if updatePerson.FirstName != "" {
				persons[i].FirstName = updatePerson.FirstName
			}
			if updatePerson.LastName != "" {
				persons[i].LastName = updatePerson.LastName
			}
			ctx.JSON(200, gin.H{"message": "Person updates"})
			return
		}
	} 
	ctx.JSON(http.StatusNotFound, gin.H{"message": "Person not found"})
}

type Person struct {
	ID string `json:"id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}

var persons = []Person{
	{ID: "1", FirstName: "Animesh", LastName: "Roy"},
	{ID: "2", FirstName: "Bikram", LastName: "Roy"},
	{ID: "3", FirstName: "Chirag", LastName: "Agarwalla"},
}