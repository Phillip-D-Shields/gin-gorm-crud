package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phillip-d-shields/gin-gorm-crud/models"
	"github.com/phillip-d-shields/gin-gorm-crud/services"
)

type MaterialController struct {
	service *services.MaterialService
}

func NewMaterialController(service *services.MaterialService) *MaterialController {
	return &MaterialController{service: service}
}

func (mc *MaterialController) GetAll(c *gin.Context) {
	materials, err := mc.service.GetAll()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, materials)
}

func (mc *MaterialController) GetByID(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	material, err := mc.service.GetByID(intID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, material)
}

func (mc *MaterialController) Create(c *gin.Context) {
	var material models.Material
	err := c.BindJSON(&material)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err = mc.service.Create(&material)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, material)
}

func (mc *MaterialController) Update(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	var material models.Material
	err = c.BindJSON(&material)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	material.ID = intID
	err = mc.service.Update(&material)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, material)
}

func (mc *MaterialController) Delete(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	err = mc.service.Delete(intID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(204, nil)
}
