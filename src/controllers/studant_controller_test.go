package controllers

import (
	"aqaryint/src/services"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCreateStudent(t *testing.T) {
	service := &services.StudantService{}
	controller := NewStudentController(service)

	// Create a new Gin HTTP request recorder
	recorder := httptest.NewRecorder()

	// Create a new Gin context using the recorder and a dummy HTTP request
	// with a JSON payload representing the new student
	router := gin.Default()
	router.POST("/students", controller.CreateStudent)
	payload := `{"name":"John Doe","age":25}`
	req, _ := http.NewRequest("POST", "/students", strings.NewReader(payload))
	router.ServeHTTP(recorder, req)

	// Assert the HTTP response status code is 201 Created
	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected status code %d but got %d", http.StatusCreated, recorder.Code)
	}

	// Assert the response body contains the created student data
	expectedResponse := `{"name":"John Doe","age":25}`
	if recorder.Body.String() != expectedResponse {
		t.Errorf("Expected response body %s but got %s", expectedResponse, recorder.Body.String())
	}
}

// Similarly, you can write tests for the GetStudent, UpdateStudent, and DeleteStudent methods
