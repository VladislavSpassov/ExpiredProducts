package models

type RecipyProduct struct {
	ID       uint   `json:"id" gorm:"primary_key" gorm:"many2many"`
	Name     string `json:"name"  `
	Quantity uint   `json:"quantity"`
}
