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

	expiry_date, _ := reader.ReadString('\n')
	expiry_date = strings.Replace(expiry_date, "\n", "", -1)
	fmt.Println(expiry_date)
	// if !IsDateValid(expiry_date) {
	// 	panic("Date is not valid")
	// }
	currentProduct.ExpiryDate = expiry_date

	fmt.Println("Change Product's quantity :")
	reader = bufio.NewReader(os.Stdin)

	quantity, _ := reader.ReadString('\n')
	quantity = strings.Replace(quantity, "\n", "", -1)
	fmt.Println(quantity)

	quantity_int, _ := strconv.ParseUint(quantity, 10, 64)
	currentProduct.Quantity = uint(quantity_int)

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
