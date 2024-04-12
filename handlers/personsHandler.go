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

// @Summary List all persons
// @Description get all persons
// @ID list-persons
// @Accept json
// @Produce json
// @Success 200 {array} models.Person
// @Router /persons [get]
func ListPersonsHandler(c *gin.Context) {
	c.JSON(200, persons)
}

// @Summary Get person details
// @Description get person details by ID
// @ID get-person-by-id
// @Accept json
// @Produce json
// @Param id path string true "Person ID"
// @Success 200 {object} models.Person
// @Failure 404 {object} map[string]string
// @Router /persons/{id} [get]
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

// @Summary Create a new person
// @Description Add a new person to the list
// @ID create-person
// @Accept json
// @Produce json
// @Param person body models.Person true "Person object to be added"
// @Success 201 {object} models.Person
// @Failure 400 {object} map[string]string
// @Router /persons [post]
func CreatePersonsHandler(c *gin.Context) {
	var person models.Person

	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	persons = append(persons, person)
	c.JSON(http.StatusCreated, gin.H{"message": "Person details created"})
}

// @Summary Delete a person
// @Description Delete a person by ID
// @ID delete-person
// @Accept json
// @Produce json
// @Param id path string true "Person ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /persons/{id} [delete]
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

// @Summary Update a person
// @Description Update a person's details by ID
// @ID update-person
// @Accept json
// @Produce json
// @Param id path string true "Person ID"
// @Param person body models.Person true "Person object with updated details"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /persons/{id} [put]
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
