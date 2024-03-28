package domain

import "time"

type Transaction struct {
	Id            int64     `json:"id" gorm:"primaryKey"`
	PayerRampId   string     `json:"payer_ramp_id" gorm:"column:payer_ramp_id"`
	PayeeRampId   string     `json:"payee_ramp_id" gorm:"column:payee_ramp_id"`
	PaymentAmount float32   `json:"payment_amount"`
	PaymentStatus string   `json:"payment_status"`
	OrderId       string    `json:"order_id" gorm:"not null;unique"`
	PaymentDate   time.Time `json:"payment_date"`
}

type UserWallet struct {
	FullName string
	RampId string
	WalletBalance float32
}

type Users struct {
	Id       int64  `json:"id" gorm:"primaryKey"`
	FullName string `json:"full_name" validate:"required,min=2,max=50"`
	RampId   string `json:"ramp_id" gorm:"unique;not null;default:null"`
	Email    string `json:"email" gorm:"not null;unique" validate:"email,required"`
	Password string `json:"password"`
}


type Wallet struct {
	Id            int64   `json:"id" gorm:"primaryKey"`
	UserId        int64   `json:"user_id"`
	WalletBalance float64 `json:"wallet_balance"`
}


