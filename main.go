package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/phillip-d-shields/gin-gorm-crud/config"
	"github.com/phillip-d-shields/gin-gorm-crud/controllers"
	"github.com/phillip-d-shields/gin-gorm-crud/models"
	"github.com/phillip-d-shields/gin-gorm-crud/routes"
	"github.com/phillip-d-shields/gin-gorm-crud/services"
)

func main() {
	db, err := config.NewDatabase()
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	// migrate models
	err = db.DB.AutoMigrate(&models.Contact{}, &models.Company{})
	if err != nil {
		log.Fatal("failed to migrate models: ", err)
	}

	// init services
	contactService := services.NewContactService(db.DB)
	companyService := services.NewCompanyService(db.DB)

	// init controllers
	contactController := controllers.NewContactController(contactService)
	companyController := controllers.NewCompanyController(companyService)

	router := gin.Default()

	// setup routes
	routes.SetupContactRoutes(router, contactController)
	routes.SetupCompanyRoutes(router, companyController)

	router.Run(":8080")
}
