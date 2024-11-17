package main

import (
	"github.com/gin-gonic/gin"
	"github.com/phillip-d-shields/gin-gorm-crud/services"
)

func main() {

	services.InitDB()
	services.SeedMockData()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "kia ora",
		})
	})

	r.Run(":8080")
}
