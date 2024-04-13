package storage

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"main.go/models"
)

var db *gorm.DB

func LoadDatabase() {

	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Person{})

}
