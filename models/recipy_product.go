package models

type RecipyProduct struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Quantity uint   `json:"quantity"`
}
