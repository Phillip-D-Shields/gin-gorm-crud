package services

import (
	"github.com/phillip-d-shields/gin-gorm-crud/models"
	"gorm.io/gorm"
)

type CompanyService struct {
	db *gorm.DB
}

func NewCompanyService(db *gorm.DB) *CompanyService {
	return &CompanyService{db: db}
}

func (cs *CompanyService) GetAll() ([]models.Company, error) {
	var companies []models.Company
	result := cs.db.Find(&companies)
	return companies, result.Error
}

func (cs *CompanyService) GetByID(id int) (*models.Company, error) {
	var company models.Company
	result := cs.db.First(&company, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &company, nil
}

func (cs *CompanyService) Create(company *models.Company) error {
	return cs.db.Create(company).Error
}

func (cs *CompanyService) Update(company *models.Company) error {
	return cs.db.Save(company).Error
}

func (cs *CompanyService) Delete(id int) error {
	return cs.db.Delete(&models.Company{}, id).Error
}
