package main

import (
	"main/controllers"
	"main/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDataBase()
	r.GET("/products", controllers.FindProduct)
	r.POST("/products", controllers.CreateProduct)
	r.GET("/products/:id", controllers.FindProducts)
	r.PATCH("/products/:id", controllers.UpdateProduct)
	r.DELETE("/products/:id", controllers.DeleteBook)
	r.Run()
}
