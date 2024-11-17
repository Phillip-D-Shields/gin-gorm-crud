package services

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/phillip-d-shields/gin-gorm-crud/models"
)

var (
	DBConn *gorm.DB
)

func InitDB() {
	var err error
	DBConn, err = gorm.Open(sqlite.Open("crud.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DBConn.AutoMigrate(models.Company{})
	DBConn.AutoMigrate(models.Contact{})
}

func SeedMockData() {
	primaryContact := models.Contact{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "jd@email.com",
		Phone:     "1234567890",
		Notes:     "Primary contact",
	}
	DBConn.Create(&primaryContact)
}
