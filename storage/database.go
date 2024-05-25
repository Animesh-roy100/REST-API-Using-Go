package storage

import (
	"fmt"

	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

	// Apply migration
	// m, err := migrate.New(
	// 	"file://migrations",
	// 	"sqlite3://test.db",
	// )
	// if err != nil {
	// 	panic(err)
	// }

	// err = m.Up()
	// if err != nil && err != migrate.ErrNoChange {
	// 	panic(err)
	// }

	// fmt.Println("migration applied")

	fmt.Println("Database connected")
}

func SetDBForTest(mockPersons []models.Person) {
	// Use the sqlite.Open function to create a Dialector for SQLite
	dialector := sqlite.Open("file::memory:?cache=shared")

	// Open the database with the dialector
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic("failed to connect mock database")
	}

	db.AutoMigrate(&models.Person{})

	for _, person := range mockPersons {
		db.Create(&person)
	}

	DB = db
}
