package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phillip-d-shields/gin-gorm-crud/models"
	"github.com/phillip-d-shields/gin-gorm-crud/services"
)

type CompanyController struct {
	service *services.CompanyService
}

func NewCompanyController(service *services.CompanyService) *CompanyController {
	return &CompanyController{service: service}
}

func (cc *CompanyController) GetAll(c *gin.Context) {
	companies, err := cc.service.GetAll()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, companies)
}

func (cc *CompanyController) GetByID(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	company, err := cc.service.GetByID(intID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, company)
}

func (cc *CompanyController) Create(c *gin.Context) {
	var company models.Company
	err := c.BindJSON(&company)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err = cc.service.Create(&company)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, company)
}

func (cc *CompanyController) Update(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	var company models.Company
	err = c.BindJSON(&company)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	company.ID = uint(intID)
	err = cc.service.Update(&company)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, company)
}

func (cc *CompanyController) Delete(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	err = cc.service.Delete(intID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(204, nil)
}
