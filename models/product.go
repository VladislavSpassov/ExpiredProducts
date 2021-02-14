package models

type Product struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	Name       string `json:"name"`
	ExpiryDate string `json:"expiry_data"`
	Quantity   uint   `json:"quantity"`
}
