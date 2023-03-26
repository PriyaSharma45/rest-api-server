package model

type CardDetails struct {
	CardNumber string `json:"card_number"`
	CVV        int    `json:"ccv"`
	ExpiryDate string `json:"expiry_date"`
	Name       string `json:"name"`
	UserID     int    `json:"user_id"`
}

type BankDetails struct {
	Address string `json:"address"`
	IFSC    string `json:"ifsc"`
	Name    string `json:"name"`
}
