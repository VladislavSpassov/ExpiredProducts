package views

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"main/models"
	"net/http"
)

//GetProducts fetch products function
func GetProducts() []models.Product {
	url := "http://localhost:8080/products"
	m := make(map[string][]models.Product)
	resp, err := http.Get(url)
	if err != nil {
		panic("cannot fetch products")
	}
	err = json.NewDecoder(resp.Body).Decode(&m)

	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)

	return m["data"]
}

//GetProductsPrint printing all current products
func GetProductsPrint() {
	products := GetProducts()
	for _, product := range products {
		fmt.Println(product)
	}

}
