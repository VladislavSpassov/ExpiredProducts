package controllers

import (
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//RecipyProductInput input structure for creating recipy product
type RecipyProductInput struct {
	Name     string `json:"name"  binding:"required"`
	Quantity uint   `json:"quantity" binding:"required"`
}

// FindRecipyProducts controller for finding all recipy products
func FindRecipyProducts(c *gin.Context) {
	var products []models.RecipyProduct
	models.DB.Find(&products)

	c.JSON(http.StatusOK, gin.H{"data": products})
}

// CreateRecipyProduct controller for creating recipy product
func CreateRecipyProduct(c *gin.Context) {
	var input RecipyProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recipyProduct := models.RecipyProduct{Name: input.Name, Quantity: input.Quantity}
	models.DB.Create(&recipyProduct)

	c.JSON(http.StatusOK, gin.H{"data": recipyProduct})
}

// FindRecipyProduct controller for finding recipy product by id
func FindRecipyProduct(c *gin.Context) { // Get model if exist
	var product models.Product

	if err := models.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// UpdateRecipyProduct controller for updating recipy product by id
func UpdateRecipyProduct(c *gin.Context) {
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

// DeleteRecipyProduct controller for deleting recipy product by id
func DeleteRecipyProduct(c *gin.Context) {
	var product models.Product
	if err := models.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
