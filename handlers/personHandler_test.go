package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"main.go/handlers"
	"main.go/models"
	"main.go/storage"
)

func TestListPersonsHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a test router
	router := gin.New()
	router.GET("/persons", handlers.ListPersonsHandler)

	// Mock your database connection and data
	mockPersons := []models.Person{
		{FirstName: "John", LastName: "Doe"},
		{FirstName: "Jane", LastName: "Smith"},
	}

	// Set the mock database before the test runs
	storage.SetDBForTest(mockPersons)

	// Create a request to pass to the handler
	req, err := http.NewRequest("GET", "/persons", nil)
	assert.NoError(t, err)

	// Create a response recorder
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Decode the response
	var responsePersons []models.Person
	err = json.Unmarshal(w.Body.Bytes(), &responsePersons)
	assert.NoError(t, err)

	// Custom comparison function to ignore ID, CreatedAt, and UpdatedAt fields
	assert.Equal(t, len(mockPersons), len(responsePersons), "Length of persons should match")
	for i, mockPerson := range mockPersons {
		assert.Equal(t, mockPerson.FirstName, responsePersons[i].FirstName, "First name should match")
		assert.Equal(t, mockPerson.LastName, responsePersons[i].LastName, "Last name should match")
	}
}

func TestGetPersonDetails(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect mock database")
	}
	storage.DB = db

	db.AutoMigrate(&models.Person{})

	router := gin.New()
	router.GET("/persons/:id", handlers.GetPersonDetails)

	mockPerson := models.Person{
		FirstName: "John",
		LastName:  "Doe",
	}
	db.Create(&mockPerson)

	req, err := http.NewRequest("GET", "/persons/"+fmt.Sprintf("%d", mockPerson.ID), nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var responsePerson models.Person
	err = json.Unmarshal(w.Body.Bytes(), &responsePerson)
	assert.NoError(t, err)

	assert.Equal(t, mockPerson.FirstName, responsePerson.FirstName, "First name should match")
	assert.Equal(t, mockPerson.LastName, responsePerson.LastName, "Last name should match")
}

func TestCreatePersonsHandler(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect mock database")
	}
	storage.DB = db

	db.AutoMigrate(&models.Person{})

	router := gin.New()
	router.POST("/persons", handlers.CreatePersonsHandler)

	personJSON := `{"firstName": "John", "lastName": "Doe"}`
	reqBody := bytes.NewBuffer([]byte(personJSON))

	req, err := http.NewRequest("POST", "/persons", reqBody)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "Person details created", response["message"], "Response message should match")
	assert.NotNil(t, response["data"], "Data should not be nil")
	data := response["data"].(map[string]interface{})
	assert.Equal(t, "John", data["firstName"], "First name should match")
	assert.Equal(t, "Doe", data["lastName"], "Last name should match")
}

func TestDeletePersonsHandler(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect mock database")
	}
	storage.DB = db

	db.AutoMigrate(&models.Person{})

	router := gin.New()
	router.DELETE("/persons/:id", handlers.DeletePersonsHandler)

	mockPerson := models.Person{
		FirstName: "John",
		LastName:  "Doe",
	}
	db.Create(&mockPerson)

	req, err := http.NewRequest("DELETE", "/persons/"+fmt.Sprintf("%d", mockPerson.ID), nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "Person removed", response["message"], "Response message should match")
}

func TestUpdatePersonHandler(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect mock database")
	}
	storage.DB = db

	db.AutoMigrate(&models.Person{})

	router := gin.New()
	router.PUT("/persons/:id", handlers.UpdatePersonHandler)

	mockPerson := models.Person{
		FirstName: "John",
		LastName:  "Doe",
	}
	db.Create(&mockPerson)

	updatePersonJSON := `{"firstName": "Jane", "lastName": "Smith"}`
	reqBody := bytes.NewBuffer([]byte(updatePersonJSON))

	req, err := http.NewRequest("PUT", "/persons/"+fmt.Sprintf("%d", mockPerson.ID), reqBody)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "Person updates", response["message"], "Response message should match")
	assert.NotNil(t, response["data"], "Data should not be nil")
	data := response["data"].(map[string]interface{})
	assert.Equal(t, "Jane", data["firstName"], "First name should be updated")
	assert.Equal(t, "Smith", data["lastName"], "Last name should be updated")
}
