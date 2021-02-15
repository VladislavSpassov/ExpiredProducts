package models

type RecipyProduct struct {
	ID       uint   `json:"id" gorm:"primary_key many2many:recipy_products" `
	Name     string `json:"name"`
	Quantity uint   `json:"quantity"`
}
