package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/models"
	"main.go/storage"
)

// @Summary List all persons
// @Description get all persons
// @Tags person
// @ID list-persons
// @Accept json
// @Produce json
// @Success 200 {array} models.Person
// @Router /persons [get]
func ListPersonsHandler(c *gin.Context) {
	var persons []models.Person
	if err := storage.DB.Find(&persons).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching persons"})
		return
	}
	c.JSON(200, persons)
}

// @Summary Get person details
// @Description get person details by ID
// @Tags person
// @ID get-person-by-id
// @Accept json
// @Produce json
// @Param id path string true "Person ID"
// @Success 200 {object} models.Person
// @Failure 404 {object} map[string]string
// @Router /persons/{id} [get]
func GetPersonDetails(c *gin.Context) {
	id := c.Param("id")

	var person models.Person
	result := storage.DB.First(&person, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "Person not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching person details"})
		}
		return
	}

	c.JSON(200, person)
}

// @Summary Create a new person
// @Description Add a new person to the list
// @Tags person
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

	result := storage.DB.Create(&person)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating person"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Person details created", "data": person})
}

// @Summary Delete a person
// @Description Delete a person by ID
// @Tags person
// @ID delete-person
// @Accept json
// @Produce json
// @Param id path string true "Person ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /persons/{id} [delete]
func DeletePersonsHandler(c *gin.Context) {
	id := c.Param("id")

	result := storage.DB.Delete(&models.Person{}, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "Person not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting person"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Person removed"})
}

// @Summary Update a person
// @Description Update a person's details by ID
// @Tags person
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

	var person models.Person
	if result := storage.DB.First(&person, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "Person not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching person"})
		}
		return
	}

	if updatePerson.FirstName != "" {
		person.FirstName = updatePerson.FirstName
	}
	if updatePerson.LastName != "" {
		person.LastName = updatePerson.LastName
	}

	result := storage.DB.Save(&person)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating person"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Person updates", "data": person})
}
