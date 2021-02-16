package controllers

import (
	"net/http"

	"main/models"

	"github.com/gin-gonic/gin"
)

//CreateProductInput input structure for creating product
type CreateProductInput struct {
	Name       string `json:"name" binding:"required"`
	ExpiryDate string `json:"expiry_date" binding:"required"`
	Quantity   uint   `json:"quantity" binding:"required"`
}

//UpdateProductInput  structure for updating product
type UpdateProductInput struct {
	Name       string `json:"name"`
	ExpiryDate string `json:"expiry_date"`
	Quantity   uint   `json:"quantity"`
}

// FindProducts controller for finding all products
func FindProducts(c *gin.Context) {
	var products []models.Product
	models.DB.Find(&products)

	c.JSON(http.StatusOK, gin.H{"data": products})
}

// CreateProduct controller for creating product
func CreateProduct(c *gin.Context) {
	var input CreateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := models.Product{Name: input.Name, ExpiryDate: input.ExpiryDate, Quantity: input.Quantity}
	models.DB.Create(&product)

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// FindProduct controller for finding product by id
func FindProduct(c *gin.Context) { // Get model if exist
	var product models.Product

	if err := models.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// UpdateProduct controller for updating product by id
func UpdateProduct(c *gin.Context) {
	var product models.Product
	if err := models.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&product).Update("expiry_date", input.ExpiryDate)
	models.DB.Model(&product).Update("quantity", input.Quantity)

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// DeleteProduct controller for deleting product by id
func DeleteProduct(c *gin.Context) {
	var product models.Product
	if err := models.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
