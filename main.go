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
	err = db.DB.AutoMigrate(&models.Contact{}, &models.Company{}, &models.Material{})
	if err != nil {
		log.Fatal("failed to migrate models: ", err)
	}

	// init services
	contactService := services.NewContactService(db.DB)
	companyService := services.NewCompanyService(db.DB)
	materialService := services.NewMaterialService(db.DB)

	// init controllers
	contactController := controllers.NewContactController(contactService)
	companyController := controllers.NewCompanyController(companyService)
	materialController := controllers.NewMaterialController(materialService)

	router := gin.Default()

	// setup routes
	routes.SetupContactRoutes(router, contactController)
	routes.SetupCompanyRoutes(router, companyController)
	routes.SetupMaterialRoutes(router, materialController)

	router.Run(":8080")
}

// TODO 1. finish service data model
// TODO 2. finish job data model
// TODO 3. slim down docker container size using multi-stage builds
