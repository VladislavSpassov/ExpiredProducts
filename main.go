package main

import (
	"bufio"
	"fmt"
	"os"

	"main/controllers"
	"main/models"
	"main/views"

	"github.com/gin-gonic/gin"
)

func setRouting(r *gin.Engine) {
	r.GET("/products", controllers.FindProducts)
	r.POST("/products", controllers.CreateProduct)
	r.GET("/products/:id", controllers.FindProduct)
	r.PATCH("/products/:id", controllers.UpdateProduct)
	r.DELETE("/products/:id", controllers.DeleteProduct)
	r.GET("/recipies", controllers.FindRecipies)
	r.POST("/recipies", controllers.CreateRecipy)
	r.PATCH("/recipies/:id", controllers.UpdateRecipy)
	r.DELETE("/recipies/:id", controllers.DeleteRecipy)
	r.Run()

}

func main() {
	models.ConnectDataBase()
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	go setRouting(r)

	reader := bufio.NewReader(os.Stdin)
	views.DisplayMenu()

	for {
		char, _, _ := reader.ReadRune()
		switch char {
		case '1':
			fmt.Println("Exitting")
			os.Exit(1)
		case '2':
			fmt.Println("Get products")
			views.GetProducts()
		case '3':
			fmt.Println("Create product")
			views.CreateProduct()
		}

	}
}
