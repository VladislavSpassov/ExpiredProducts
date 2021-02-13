package models

import "time"

type Product struct {
	ID         uint      `json:"id" gorm:"primary_key"`
	Name       string    `json:"title"`
	ExpiryDate time.Time `json:"expiry_data"`
	Quantity   uint      `json:"quantity"`
}
