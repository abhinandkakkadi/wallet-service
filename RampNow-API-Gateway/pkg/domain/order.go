package domain

type Transaction struct {
	PayerRampId string `json:"payer_ramp_id"`
	PayeeRampId string `json:"payee_ramp_id"`
	PaymentAmount float64 `json:"payment_amount"`
}

type Order struct {
	ID           string  `json:"id" swaggerignore:"true"`
	Status       string  `json:"status"`
	Item         []Item  `json:"items"`
	Total        float64 `json:"total"`
	CurrencyUnit string  `json:"currency_unit"`
}

type Item struct {
	ID          string  `json:"id" `
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}
type UpdateOrder struct {
	ID     string `json:"id" `
	Status string `json:"status"`
}
