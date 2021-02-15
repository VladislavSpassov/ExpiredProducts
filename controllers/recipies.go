package controllers

import (
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateRecipyInput struct {
	Name           string                 `json:"name" binding:"required"`
	RecipyProducts []models.RecipyProduct `json:"recipy_products" binding:"required"`
}

type UpdateRecipyInput struct {
	Name           string                 `json:"name"`
	RecipyProducts []models.RecipyProduct `json:"recipy_products"`
}

func FindRecipies(c *gin.Context) {
	var recipies []models.Recipy
	models.DB.Preload("RecipyProducts").Find(&recipies)

	c.JSON(http.StatusOK, gin.H{"data": recipies})
}

func CreateRecipy(c *gin.Context) {
	var input CreateRecipyInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for recipy_product := range input.RecipyProducts {
		models.DB.Save(&recipy_product)
	}

	recipy := models.Recipy{RecipyProducts: input.RecipyProducts, Name: input.Name}
	models.DB.Create(&recipy)

	c.JSON(http.StatusOK, gin.H{"data": recipy})
}

func DeleteRecipy(c *gin.Context) {
	var recipy models.Recipy
	if err := models.DB.Where("id = ?", c.Param("id")).First(&recipy).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&recipy)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func UpdateRecipy(c *gin.Context) {

	var recipy models.Recipy
	if err := models.DB.Where("id = ?", c.Param("id")).First(&recipy).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateRecipyInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&recipy).Updates(models.Recipy{Name: input.Name, RecipyProducts: input.RecipyProducts})

	c.JSON(http.StatusOK, gin.H{"data": recipy})
}
