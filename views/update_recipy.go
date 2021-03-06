package views

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"main/models"
	"net/http"
	"os"
	"strings"
)

//UpdateRecipy Functiion to update recipy by a given name
func UpdateRecipy() {

	fmt.Println("What recipy do you want to update?:")
	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	recipies := GetRecipies()
	var currentRecipy models.Recipy
	var isFound bool = false
	for _, recipy := range recipies {
		if recipy.Name == text {
			currentRecipy = recipy
			isFound = true
		}
	}
	if !isFound {
		panic("Not found recipy")
	}
	fmt.Println("Change Recipy's expiry date:")
	reader = bufio.NewReader(os.Stdin)

	fmt.Println("Change Product's quantity :")
	reader = bufio.NewReader(os.Stdin)

	quantity, _ := reader.ReadString('\n')
	quantity = strings.Replace(quantity, "\n", "", -1)
	fmt.Println(quantity)

	json, err := json.Marshal(currentRecipy)

	url := "http://localhost:8080/products/" + fmt.Sprint(currentRecipy.ID)
	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(json))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic("Cannot update product!")
	}
	defer resp.Body.Close()

}
