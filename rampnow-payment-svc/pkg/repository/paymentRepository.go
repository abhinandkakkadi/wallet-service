package repository

import (
	"context"
	"errors"

	"github.com/abhinandkakkadi/rampnow-payment-svc/pkg/domain"
	interfaces "github.com/abhinandkakkadi/rampnow-payment-svc/pkg/repository/interface"
	"gorm.io/gorm"
)

type paymentDatabase struct {
	DB *gorm.DB
}

func (o *paymentDatabase) GetUserIdFromRampId(ctx context.Context, rampId string) (int, error) {
	var userID int 

	err := o.DB.Raw("SELECT id FROM users WHERE ramp_id = ?", rampId).Scan(&userID).Error
	if err != nil {
		return 0, errors.New("transaction failed due to server inconsistencies")
	}

	return userID, nil
}

func (o *paymentDatabase) GetWalletAmountFromUserId(ctx context.Context, userId int) (float64, error) {
	var walletAmount float64 

	err := o.DB.Raw("SELECT wallet_balance FROM wallets WHERE  user_id = ?", userId).Scan(&walletAmount).Error
	if err != nil {
		return 0, errors.New("transaction failed due to server inconsistencies")
	}

	return walletAmount, nil
}

func (o *paymentDatabase) IncrementWalletBalance(ctx context.Context, userId int, paymentAmount float64) (error) {
	err := o.DB.Exec("UPDATE wallets SET wallet_balance = wallet_balance + ? WHERE user_id = ?", paymentAmount, userId).Error
	return err
}

func (o *paymentDatabase) DecrementWalletBalance(ctx context.Context, userId int, paymentAmount float64) (error) {
	err := o.DB.Exec("UPDATE wallets SET wallet_balance = wallet_balance - ? WHERE user_id = ?", paymentAmount, userId).Error
	return err
}

func (o *paymentDatabase) FindUser(ctx context.Context, rampId string) (int, error) {
	var exists int

	err := o.DB.Raw("SELECT COUNT(*) FROM users WHERE ramp_id = ?", rampId).Scan(&exists).Error
	if err != nil {
		return 0, errors.New("transaction failed due to server inconsistencies")
	}

	return exists, nil
}

func (o *paymentDatabase) FindByID(ctx context.Context, userId int) (domain.Users, error) {
	var user domain.Users
	err := o.DB.First(&user, userId).Error

	return user, err
}

func (o *paymentDatabase) GetAllTransactions(ctx context.Context) ([]domain.Transaction, error) {
	var transaction []domain.Transaction
	err := o.DB.Find(&transaction).Error

	return transaction, err
}



func (o *paymentDatabase) CreateTransaction(ctx context.Context, transaction domain.Transaction) (domain.Transaction, error) {
	err := o.DB.Save(&transaction).Error
	return transaction, err
}


func NewPaymentRepository(DB *gorm.DB) interfaces.PaymentRepository {
	return &paymentDatabase{DB}
}
