package controllers

import (
	"net/http"
	"time"

	"main/models"

	"github.com/gin-gonic/gin"
)

type CreateProductInput struct {
	Name       string    `json:"name" binding:"required"`
	ExpiryDate time.Time `json:"expiry_date" binding:"required"`
}

type UpdateProductInput struct {
	Name       string    `json:"name"`
	ExpiryDate time.Time `json:"expiry_date"`
}

func FindProducts(c *gin.Context) {
	var products []models.Product
	models.DB.Find(&products)

	c.JSON(http.StatusOK, gin.H{"data": products})
}

func CreateProduct(c *gin.Context) {
	// Validate input
	var input CreateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	product := models.Product{Name: input.Name, ExpiryDate: input.ExpiryDate}
	models.DB.Create(&product)

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func FindProduct(c *gin.Context) { // Get model if exist
	var product models.Product

	if err := models.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func UpdateProduct(c *gin.Context) {
	// Get model if exist
	var product models.Product
	if err := models.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&product).Updates(models.Product{Name: input.Name, ExpiryDate: input.ExpiryDate})

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func DeleteBook(c *gin.Context) {
	var product models.Product
	if err := models.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func GetAllValidProducts(c *gin.Context) {

}
