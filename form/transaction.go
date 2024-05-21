package form

type Transaction struct {
	ID uint64 `json:"transaction_id"`
	Date string `json:"date"`
	Amount float64 `json:"amount"`
	TransactionType uint64 `json:"transaction_type"`
	Note string `json:"note"`
	ImageURL string `json:"image_url"`
	SpenderID uint64 `json:"spender_id"`
}
