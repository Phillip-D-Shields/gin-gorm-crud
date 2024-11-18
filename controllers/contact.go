package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phillip-d-shields/gin-gorm-crud/models"
	"github.com/phillip-d-shields/gin-gorm-crud/services"
)

type ContactController struct {
	service *services.ContactService
}

func NewContactController(service *services.ContactService) *ContactController {
	return &ContactController{service: service}
}

func (cc *ContactController) GetAll(c *gin.Context) {
	contacts, err := cc.service.GetAll()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, contacts)
}

func (cc *ContactController) GetByID(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	contact, err := cc.service.GetByID(intID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, contact)
}

func (cc *ContactController) Create(c *gin.Context) {
	var contact models.Contact
	err := c.BindJSON(&contact)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err = cc.service.Create(&contact)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, contact)
}

func (cc *ContactController) Update(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	var contact models.Contact
	err = c.BindJSON(&contact)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	contact.ID = uint(intID)
	err = cc.service.Update(&contact)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, contact)
}

func (cc *ContactController) Delete(c *gin.Context) {
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
