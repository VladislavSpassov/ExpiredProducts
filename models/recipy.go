package models

import "time"

type Recipy struct {
	ID         uint      `json:"id" gorm:"primary_key"`
	Products   []Product `json:"products"`
	ExpiryDate time.Time `json:"expiry_data"`
}
