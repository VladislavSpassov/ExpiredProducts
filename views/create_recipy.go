package views

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"main/models"
	"net/http"
	"os"
	"strconv"
	"strings"
)

//RecipyProductInput structure for create recipy product
type RecipyProductInput struct {
	Name     string `json:"name"  binding:"required"`
	Quantity uint   `json:"quantity" binding:"required"`
}

//CreateRecipy function for creating recipy with given name and recipy products by user
func CreateRecipy() {
	url := "http://localhost:8080/recipies"

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Recipy name: ")
	name, err := reader.ReadString('\n')
	name = strings.TrimSuffix(name, "\n")

	if err != nil {
		panic("Not correct name")
	}

	var recipyProducts []models.RecipyProduct

	for {
		fmt.Println("Do you want to add product to recipy?")
		reader = bufio.NewReader(os.Stdin)
		ans, _, _ := reader.ReadRune()
		if ans == 'n' {
			break
		}

		var product models.RecipyProduct

		fmt.Println("Add product's name")
		reader := bufio.NewReader(os.Stdin)

		productName, _ := reader.ReadString('\n')
		productName = strings.TrimSuffix(productName, "\n")
		fmt.Println("Add product's quantity")
		reader = bufio.NewReader(os.Stdin)

		quantity, _ := reader.ReadString('\n')
		quantity = strings.TrimSuffix(quantity, "\n")

		quantityInt, _ := strconv.ParseUint(quantity, 10, 64)

		product.Name = productName
		product.Quantity = uint(quantityInt)
		recipyProducts = append(recipyProducts, product)
		fmt.Println(productName, quantityInt)
	}
	postBody, _ := json.Marshal(map[string]interface{}{
		"name":            name,
		"recipy_products": recipyProducts,
	})
	responseBody := bytes.NewBuffer(postBody)
	_, err = http.Post(url, "application/json", responseBody)

	if err != nil {
		panic("cannot create recipy")
	}
}
