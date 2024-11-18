package routes

import (
	"github.com/gin-gonic/gin"

	. "github.com/phillip-d-shields/gin-gorm-crud/controllers"
)

func SetupContactRoutes(router *gin.Engine, cc *ContactController) {

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	v1 := router.Group("/api/v1")
	{
		// contacts
		v1.GET("/contacts", cc.GetAll)
		v1.GET("/contacts/:id", cc.GetByID)
		v1.POST("/contacts", cc.Create)
		v1.PUT("/contacts/:id", cc.Update)
		v1.DELETE("/contacts/:id", cc.Delete)
	}
}

func SetupCompanyRoutes(router *gin.Engine, cc *CompanyController) {

	v1 := router.Group("/api/v1")
	{
		// companies
		v1.GET("/companies", cc.GetAll)
		v1.GET("/companies/:id", cc.GetByID)
		v1.POST("/companies", cc.Create)
		v1.PUT("/companies/:id", cc.Update)
		v1.DELETE("/companies/:id", cc.Delete)
	}
}

func SetupMaterialRoutes(router *gin.Engine, mc *MaterialController) {

	v1 := router.Group("/api/v1")
	{
		// materials
		v1.GET("/materials", mc.GetAll)
		v1.GET("/materials/:id", mc.GetByID)
		v1.POST("/materials", mc.Create)
		v1.PUT("/materials/:id", mc.Update)
		v1.DELETE("/materials/:id", mc.Delete)
	}
}
