package interfaces

import (
	"context"

	"github.com/abhinandkakkadi/rampnow-payment-svc/pkg/domain"
)

type PaymentRepository interface {
	FindByID(ctx context.Context, userId int) (domain.Users, error)
	GetAllTransactions(ctx context.Context) ([]domain.Transaction, error)
	CreateTransaction(ctx context.Context, order domain.Transaction) (domain.Transaction, error)
	FindUser(ctx context.Context, rampId string) (int, error)
	GetUserIdFromRampId(ctx context.Context, rampId string) (int, error)
	GetWalletAmountFromUserId(ctx context.Context, userId int) (float64, error)
	IncrementWalletBalance(ctx context.Context, userId int, paymentAmount float64) (error)
	DecrementWalletBalance(ctx context.Context, userId int, paymentAmount float64) (error)
}
