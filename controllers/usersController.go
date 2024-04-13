package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"main.go/models"
	"main.go/storage"
)

func Signup(c *gin.Context) {
	// Get the email/pass off req Body

	var user models.User

	if c.ShouldBindJSON(&user) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password.",
		})
		return
	}

	// Create the user
	newUser := models.User{Email: user.Email, Password: string(hash)}

	result := storage.DBUser.Create(&newUser)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user.",
		})
	}

	// Respond
	fmt.Println("user created")
	c.JSON(http.StatusOK, gin.H{"message": "user created", "user": newUser})
}
