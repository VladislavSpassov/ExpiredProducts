package views

import (
	"fmt"
	"io/ioutil"
)

//DisplayMenu function to display options' menu to user
func DisplayMenu() {
	content, err := ioutil.ReadFile("./views/display_menu.txt")

	if err != nil {
		panic("Failed to open the file!")
	}

	fmt.Println(string(content))

}
