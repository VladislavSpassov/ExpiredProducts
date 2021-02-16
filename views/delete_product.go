package views

import (
	"bufio"
	"fmt"
	"main/models"
	"net/http"
	"os"
	"strings"
)

//DeleteProduct function to delete a product by given product's name by user
func DeleteProduct() {
	fmt.Println("What product do you want to delete?:")
	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	products := GetProducts()
	var currentProduct *models.Product
	for _, product := range products {
		if product.Name == text {
			currentProduct = &product
		}
	}
	if currentProduct == nil {
		panic("No such product is found to delete!")
	}

	// Create request
	req, _ := http.NewRequest(http.MethodDelete, "http://localhost:8080/product/:id", nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic("Cannot delete product!")
	}
	defer resp.Body.Close()

}
