package views

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

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
	quantity, err := reader.ReadString('\n')
	quantity = strings.TrimSuffix(quantity, "\n")

	if err != nil {
		panic("Not correct quantity")
	}

	postBody, _ := json.Marshal(map[string]string{
		"name":        name,
		"expiry_date": date,
		"quantity":    quantity,
	})
	responseBody := bytes.NewBuffer(postBody)
	fmt.Println(responseBody)
	_, err = http.Post(url, "application/json", responseBody)

	if err != nil {
		panic("cannot create product")
	}
}
