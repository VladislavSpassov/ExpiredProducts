package views

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateRecipyProduct(product_name string, quantity string) {
	url := "http://localhost:8080/recipy_products"

	// reader := bufio.NewReader(os.Stdin)

	// fmt.Println("Product name: ")
	// name, err := reader.ReadString('\n')
	// name = strings.TrimSuffix(name, "\n")

	// if err != nil {
	// 	panic("Not correct name")
	// }
	// fmt.Println("Expiry date in format YYYY-MM-DD: ")

	// date, err := reader.ReadString('\n')
	// date = strings.TrimSuffix(date, "\n")

	// if err != nil {
	// 	panic("Not correct expiry date")
	// }
	// const layout = "2006-01-02T15:04:05.000Z"

	// if !IsDateValid(date) {
	// 	panic("Incorect date!")
	// }

	// fmt.Println("Product's quantity: ")
	// quantity, err := reader.ReadString('\n')
	// quantity = strings.TrimSuffix(quantity, "\n")

	// if err != nil {
	// 	panic("Not correct quantity")
	// }
	postBody, _ := json.Marshal(map[string]string{
		"name":     product_name,
		"quantity": quantity,
	})
	responseBody := bytes.NewBuffer(postBody)
	fmt.Println(responseBody)
	_, err := http.Post(url, "application/json", responseBody)

	if err != nil {
		panic("cannot create product")
	}
}
