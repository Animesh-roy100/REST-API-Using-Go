package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/models"
)

var persons = []models.Person{
	{ID: "1", FirstName: "Animesh", LastName: "Roy"},
	{ID: "2", FirstName: "Bikram", LastName: "Roy"},
	{ID: "3", FirstName: "Chirag", LastName: "Agarwalla"},
}

func ListPersonsHandler(c *gin.Context) {
	c.JSON(200, persons)
}

func GetPersonDetails(c *gin.Context) {
	id := c.Param(("id"))

	for _, val := range persons {
		if val.ID == id {
			c.JSON(200, val)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Person not found"})
}

func CreatePersonsHandler(c *gin.Context) {
	var person models.Person

	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	persons = append(persons, person)
	c.JSON(http.StatusCreated, gin.H{"message": "Person details created"})
}

func DeletePersonsHandler(c *gin.Context) {
	id := c.Param("id")

	for i, a := range persons {
		if a.ID == id {
			persons = append(persons[:i], persons[i+1:]...)
			c.JSON(200, gin.H{"message": "Person removed"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"Message": "Person not found"})
}

func UpdatePersonHandler(c *gin.Context) {
	id := c.Param("id")

	var updatePerson models.Person

	if err := c.ShouldBindJSON(&updatePerson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
			c.JSON(200, gin.H{"message": "Person updates"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Person not found"})
}
