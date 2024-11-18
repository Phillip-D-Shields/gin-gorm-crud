package services

import (
	"gorm.io/gorm"

	"github.com/phillip-d-shields/gin-gorm-crud/models"
)

type ContactService struct {
	db *gorm.DB
}

func NewContactService(db *gorm.DB) *ContactService {
	return &ContactService{db: db}
}

func (cs *ContactService) GetAll() ([]models.Contact, error) {
	var contacts []models.Contact
	result := cs.db.Find(&contacts)
	return contacts, result.Error
}

func (cs *ContactService) GetByID(id int) (*models.Contact, error) {
	var contact models.Contact
	result := cs.db.First(&contact, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &contact, nil
}

func (cs *ContactService) Create(contact *models.Contact) error {
	return cs.db.Create(contact).Error
}

func (cs *ContactService) Update(contact *models.Contact) error {
	return cs.db.Save(contact).Error
}

func (cs *ContactService) Delete(id int) error {
	return cs.db.Delete(&models.Contact{}, id).Error
}
