package views

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"main/models"
	"net/http"
)

type GetProductStruct struct {
	ID         uint
	Name       string
	ExpiryDate string
	Quantity   uint
}

func GetProducts() {
	url := "http://localhost:8080/products"
	m := make(map[string][]models.Product)
	resp, err := http.Get(url)
	if err != nil {
		panic("cannot fetch products")
	}
	err = json.NewDecoder(resp.Body).Decode(&m)

	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	for _, product := range m {
		fmt.Println(product)
	}
}
