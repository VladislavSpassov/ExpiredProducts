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

//UpdateProduct function to update a product by a given name
func UpdateProduct() {

	fmt.Println("What product do you want to update?:")
	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	products := GetProducts()
	var currentProduct models.Product
	var isFound bool = false
	for _, product := range products {
		if product.Name == text {
			currentProduct = product
			isFound = true
		}
	}
	if !isFound {
		panic("Not found product")
	}

	fmt.Println("Change Product's expiry date:")
	reader = bufio.NewReader(os.Stdin)

	expiryDate, _ := reader.ReadString('\n')
	expiryDate = strings.Replace(expiryDate, "\n", "", -1)
	if !IsDateValid(expiryDate) {
		panic("Date is not valid")
	}
	currentProduct.ExpiryDate = expiryDate

	fmt.Println("Change Product's quantity :")
	reader = bufio.NewReader(os.Stdin)

	quantity, _ := reader.ReadString('\n')
	quantity = strings.Replace(quantity, "\n", "", -1)
	fmt.Println(quantity)

	quantityInt, _ := strconv.ParseUint(quantity, 10, 64)
	currentProduct.Quantity = uint(quantityInt)

	json, err := json.Marshal(currentProduct)

	url := "http://localhost:8080/products/" + fmt.Sprint(currentProduct.ID)
	fmt.Println(currentProduct.ExpiryDate)
	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(json))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic("Cannot update product!")
	}
	defer resp.Body.Close()

}
