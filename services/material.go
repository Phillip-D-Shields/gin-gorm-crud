package services

import (
	"gorm.io/gorm"

	"github.com/phillip-d-shields/gin-gorm-crud/models"
)

type MaterialService struct {
	db *gorm.DB
}

func NewMaterialService(db *gorm.DB) *MaterialService {
	return &MaterialService{db: db}
}

func (ms *MaterialService) GetAll() ([]models.Material, error) {
	var materials []models.Material
	result := ms.db.Find(&materials)
	return materials, result.Error
}

func (ms *MaterialService) GetByID(id int) (*models.Material, error) {
	var material models.Material
	result := ms.db.First(&material, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &material, nil
}

func (ms *MaterialService) Create(material *models.Material) error {
	return ms.db.Create(material).Error
}

func (ms *MaterialService) Update(material *models.Material) error {
	return ms.db.Save(material).Error
}

func (ms *MaterialService) Delete(id int) error {
	return ms.db.Delete(&models.Material{}, id).Error
}
