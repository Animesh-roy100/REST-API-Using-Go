package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"main.go/models"
	"main.go/storage"
)

// @Summary User Signup
// @Description Create a new user account by signing up with email and password
// @Tags user
// @Accept json
// @Produce json
// @Param user body models.User true "User Data"
// @Success 200 {object} models.User "Successfully created user"
// @Failure 400 {object} map[string]string "Invalid input"
// @Router /signup [post]
func Signup(c *gin.Context) {
	// Get the email/pass off req Body

	var user models.User

	if c.ShouldBindJSON(&user) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})

		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create the user
	newUser := models.User{Email: user.Email, Password: string(hash)}

	result := storage.DBUser.Create(&newUser)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create user"})
	}

	// Respond
	fmt.Println("user created")
	c.JSON(http.StatusOK, gin.H{"message": "user created", "user": newUser})
}

// @Summary User Login
// @Description Authenticate user by logging in with email and password, and generate JWT token for authorization
// @Tags user
// @Accept json
// @Produce json
// @Param user body models.User true "User Data"
// @Success 200 {object} map[string]string "Successfully logged in"
// @Failure 400 {object} map[string]string "Invalid email or password"
// @Router /login [post]
func Login(c *gin.Context) {
	// Get email & pass off req body
	var user models.User

	if c.ShouldBindJSON(&user) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})

		return
	}

	// Look up for requested user
	var storedUser models.User

	storage.DBUser.First(&storedUser, "email = ?", user.Email)

	if storedUser.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}

	// Compare sent in password with saved users password
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create token"})
		return
	}

	// Respond
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "login successful"})
}
