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

	var recipy_products []RecipyProductInput

	for {
		fmt.Println("Do you want to add product to recipy?")
		reader = bufio.NewReader(os.Stdin)
		ans, _, _ := reader.ReadRune()
		if ans == 'n' {
			break
		}

		var product RecipyProductInput

		fmt.Println("Add product's name")
		reader := bufio.NewReader(os.Stdin)

		product_name, _ := reader.ReadString('\n')
		product_name = strings.TrimSuffix(product_name, "\n")
		fmt.Println("Add product's quantity")
		reader = bufio.NewReader(os.Stdin)

		quantity, _ := reader.ReadString('\n')
		quantity_int, _ := strconv.ParseUint(quantity, 10, 64)

		product.Name = product_name
		product.Quantity = uint(quantity_int)
		recipy_products = append(recipy_products, product)

	}
	

	var recipy models.Recipy
	// recipy.RecipyProducts = recipy_products
	// recipy.Name = name
	postBody, _ := json.Marshal(map[string]interface{}{
		"name":            recipy.Name,
		"recipy_products": recipy.RecipyProducts,
	})
	responseBody := bytes.NewBuffer(postBody)
	fmt.Println(responseBody)
	_, err = http.Post(url, "application/json", responseBody)

	if err != nil {
		panic("cannot create recipy")
	}
}
