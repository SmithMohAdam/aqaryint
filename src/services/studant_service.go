package services

import (
	"aqaryint/src/models"
	"aqaryint/src/repositories"
)

type StudantService struct {
	repo *repositories.StudantRepository
}

func NewStudentService(repo *repositories.StudantRepository) *StudantService {
	return &StudantService{
		repo: repo,
	}
}

func (svc *StudantService) CreateStudent(studant *models.Student) error {
	return svc.repo.CreateStudant(studant)
}

func (svc *StudantService) GetStudent(id uint) (*models.Student, error) {
	return svc.repo.GetStudant(id)

}

func (svc *StudantService) UpdateStudent(studant *models.Student) error {
	return svc.repo.UpdateStudant(studant)

}

func (svc *StudantService) DeleteStudent(id uint) error {
	return svc.repo.DeleteStudant(id)
}
