package repositories

import (
	"aqaryint/src/models"

	"gorm.io/gorm"
)

type StudantRepository struct {
	db *gorm.DB
}

func NewStudantRepository(db *gorm.DB) *StudantRepository {
	return &StudantRepository{
		db: db,
	}
}

func (repo *StudantRepository) CreateStudant(studant *models.Student) error {
	result := repo.db.Create(studant)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *StudantRepository) GetStudant(id uint) (*models.Student, error) {
	var studant models.Student
	result := repo.db.First(&studant, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &studant, nil
}

func (repo *StudantRepository) UpdateStudant(studant *models.Student) error {
	result := repo.db.Save(studant)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *StudantRepository) DeleteStudant(id uint) error {
	result := repo.db.Delete(&models.Student{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
