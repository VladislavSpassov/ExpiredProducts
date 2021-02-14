package models

type Recipy struct {
	ID             uint            `json:"id" gorm:"primary_key"`
	RecipyProducts []RecipyProduct `gorm:"many2many" json:"recipy_products"`
	Name           string          `json:"name"`
}
