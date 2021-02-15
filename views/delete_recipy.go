package views

import (
	"bufio"
	"fmt"
	"main/models"
	"net/http"
	"os"
	"strings"
)

func DeleteRecipy() {
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
	req, _ := http.NewRequest(http.MethodDelete, "http://localhost:8080/recipies/:id", nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic("Cannot delete recipy!")
	}
	defer resp.Body.Close()

}
