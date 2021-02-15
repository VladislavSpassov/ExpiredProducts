package views

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"main/models"
	"net/http"
)

func GetRecipies() []models.Recipy {
	url := "http://localhost:8080/recipies"
	m := make(map[string][]models.Recipy)
	resp, err := http.Get(url)
	if err != nil {
		panic("cannot fetch recipies")
	}

	fmt.Println(models.DB.Migrator().HasTable("RecipyProduct"))
	err = json.NewDecoder(resp.Body).Decode(&m)

	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)

	return m["data"]
}
func GetRecipiesPrint() {
	products := GetRecipies()
	for _, product := range products {
		fmt.Println(product)
	}

}
