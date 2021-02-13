type CreateRecipyInput struct {
	Name       string    `json:"name" binding:"required"`
	ExpiryDate time.Time `json:"expiry_date" binding:"required"`
}
