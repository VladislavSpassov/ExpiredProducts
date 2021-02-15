package models

type Product struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	Name       string `gorm:"unique" json:"name"`
	ExpiryDate string `json:"expiry_date"`
	Quantity   uint   `json:"quantity"`
}
