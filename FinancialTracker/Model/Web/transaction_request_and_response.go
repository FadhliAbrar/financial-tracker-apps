package web

type TransactionRequest struct {
	Type        string  `json:"type"`
	Amount      float32 `json:"amount"`
	Description string  `json:"description"`
	Category_id int     `json:"category_id"`
	Date        string  `json:"date"`
}

type TransactionResponse struct {
	Id          int     `json:"id"`
	Type        string  `json:"type"`
	Amount      float32 `json:"amount"`
	Description string  `json:"description"`
	Category_id int     `json:"category_id"`
	Date        string  `json:"date"`
}

type TransactionWebResponses struct {
	Code   int
	Status string
	Data   interface{}
}
