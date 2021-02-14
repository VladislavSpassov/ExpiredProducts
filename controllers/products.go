package controllers

import (
	"net/http"
	"strconv"

	"main/models"

	"github.com/gin-gonic/gin"
)

type CreateProductInput struct {
	Name       string `json:"name" binding:"required"`
	ExpiryDate string `json:"expiry_date" binding:"required"`
	Quantity   string `json:"quantity" binding:"required"`
}

type UpdateProductInput struct {
	Name       string `json:"name"`
	ExpiryDate string `json:"expiry_date"`
}

func FindProducts(c *gin.Context) {
	var products []models.Product
	models.DB.Find(&products)

	c.JSON(http.StatusOK, gin.H{"data": products})
}

func CreateProduct(c *gin.Context) {
	var input CreateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	quantity, _ := strconv.ParseUint(input.Quantity, 10, 64)
	product := models.Product{Name: input.Name, ExpiryDate: input.ExpiryDate, Quantity: uint(quantity)}
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

func DeleteProduct(c *gin.Context) {
	var product models.Product
	if err := models.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
