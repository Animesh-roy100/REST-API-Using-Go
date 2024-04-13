package storage

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"main.go/models"
)

var DB *gorm.DB
var DBUser *gorm.DB

func LoadDatabase() {

	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect test database")
	}

	DBUser, err = gorm.Open(sqlite.Open("user.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect user database")
	}

	DB.AutoMigrate(&models.Person{})
	DBUser.AutoMigrate(&models.User{})

	fmt.Println("Database connected")
}
