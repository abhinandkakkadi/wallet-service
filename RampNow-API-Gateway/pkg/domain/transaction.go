package domain

type Transaction struct {
	PayerRampId string `json:"payer_ramp_id"`
	PayeeRampId string `json:"payee_ramp_id"`
	PaymentAmount float64 `json:"payment_amount"`
}

