package features

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Feature: Hello World
//
//	As a user
//	I want to access the hello world endpoint
//	So I can see the message "Hello World!"
//
//	Scenario: Accessing the hello world endpoint
//		Given I access the hello world endpoint
//		When I send a GET request to "/api/hello-world"
//		Then the response status code should be 200
//		And the response body should be {"message": "Hello World!"}

func TestHelloWorld(t *testing.T) {
	// Given I access the hello world endpoint
	gin.SetMode(gin.TestMode) // Set gin to test mode
	router := gin.Default()
	router.GET("/api/hello-world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	req, err := http.NewRequest("GET", "/api/hello-world", nil)
	if err != nil {
		t.Fatal(err)
	}

	// When I send a GET request to "/api/hello-world"
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req) // Use router to serve the request

	// Then the response status code should be 200
	assert.Equal(t, http.StatusOK, rr.Code)

	// And the response body should be {"message": "Hello World!"}
	var response map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "Hello World!", response["message"])
}
