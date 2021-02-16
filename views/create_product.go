package views

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// IsDateValid Check if string date is in format YYYY-MM-DD
func IsDateValid(date string) bool {
	const layout = "2006-01-02T15:04:05.000Z"

	_, err := time.Parse(layout, date+"T00:00:00.000Z")
	if err != nil {
		print(err)
		return false
	}
	return true
}

// CreateProduct Function to create product by user with given product's name, expiry date and quantity
func CreateProduct() {
	url := "http://localhost:8080/products"

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Product name: ")
	name, err := reader.ReadString('\n')
	name = strings.TrimSuffix(name, "\n")

	if err != nil {
		panic("Not correct name")
	}
	fmt.Println("Expiry date in format YYYY-MM-DD: ")

	date, err := reader.ReadString('\n')
	date = strings.TrimSuffix(date, "\n")

	if err != nil {
		panic("Not correct expiry date")
	}
	const layout = "2006-01-02T15:04:05.000Z"

	if !IsDateValid(date) {
		panic("Incorect date!")
	}

	fmt.Println("Product's quantity: ")
	quantity, err := reader.ReadString('\n')
	quantity = strings.TrimSuffix(quantity, "\n")

	quantityInt, _ := strconv.ParseUint(quantity, 10, 64)
	if err != nil {
		panic("Not correct quantity")
	}

	postBody, _ := json.Marshal(map[string]interface{}{
		"name":        name,
		"expiry_date": date,
		"quantity":    quantityInt,
	})
	responseBody := bytes.NewBuffer(postBody)
	fmt.Println(responseBody)
	_, err = http.Post(url, "application/json", responseBody)

	if err != nil {
		panic("cannot create product")
	}
}
