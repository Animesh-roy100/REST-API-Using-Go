package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"

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
