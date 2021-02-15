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

type RecipyProductInput struct {
	Name     string `json:"name"  binding:"required"`
	Quantity uint   `json:"quantity" binding:"required"`
}

func CreateRecipy() {
	url := "http://localhost:8080/recipies"

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Recipy name: ")
	name, err := reader.ReadString('\n')
	name = strings.TrimSuffix(name, "\n")

	if err != nil {
		panic("Not correct name")
	}

	var recipy_products []models.RecipyProduct

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

		product_name, _ := reader.ReadString('\n')
		product_name = strings.TrimSuffix(product_name, "\n")
		fmt.Println("Add product's quantity")
		reader = bufio.NewReader(os.Stdin)

		quantity, _ := reader.ReadString('\n')
		quantity = strings.TrimSuffix(quantity, "\n")

		quantity_int, _ := strconv.ParseUint(quantity, 10, 64)

		product.Name = product_name
		product.Quantity = uint(quantity_int)
		recipy_products = append(recipy_products, product)
		fmt.Println(product_name, quantity_int)
	}
	fmt.Println(recipy_products)
	postBody, _ := json.Marshal(map[string]interface{}{
		"name":            name,
		"recipy_products": recipy_products,
	})
	responseBody := bytes.NewBuffer(postBody)
	_, err = http.Post(url, "application/json", responseBody)

	if err != nil {
		panic("cannot create recipy")
	}
}
