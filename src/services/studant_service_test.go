package services

import (
	"aqaryint/src/models"
	"aqaryint/src/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStudantService_CreateStudent(t *testing.T) {
	repo := &repositories.StudantRepository{}
	service := NewStudentService(repo)

	studant := &models.Student{
		Name:  "mohammed",
		Age:   28,
		City:  "khartoum",
		Class: rune('A'),
		//Fit:   true,
	}

	err := service.CreateStudent(studant)
	assert.NoError(t, err)
}

func TestStudantService_GetStudent(t *testing.T) {
	repo := &repositories.StudantRepository{}
	service := NewStudentService(repo)

	studant := &models.Student{
		Name:  "mohammed",
		Age:   28,
		City:  "khartoum",
		Class: rune('A'),
		//Fit:   true,
	}

	// Create the student first
	err := service.CreateStudent(studant)
	assert.NoError(t, err)

	// Get the student by ID
	resultStudant, err := service.GetStudent(1)
	assert.NoError(t, err)
	assert.Equal(t, studant, resultStudant)
}

// Similarly, you can write tests for the remaining methods (UpdateStudent, DeleteStudent)
