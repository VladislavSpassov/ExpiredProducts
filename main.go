package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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
	r.GET("/recipy_products", controllers.FindRecipyProducts)
	r.POST("/recipy_products", controllers.CreateRecipyProduct)
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
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		switch text {
		case "1":
			fmt.Println("Exitting")
			os.Exit(1)
		case "2":
			fmt.Println("Get products")
			views.GetProductsPrint()
		case "3":
			fmt.Println("Create product")
			views.CreateProduct()
		case "4":
			fmt.Println("Update product")
			views.UpdateProduct()
		case "5":
			fmt.Println("Delete product")
			views.DeleteProduct()
		case "7":
			fmt.Println("Get recipies")
			views.GetRecipiesPrint()
		case "8":
			fmt.Println("Add recipy")
			views.CreateRecipy()
		}

	}
}
