package interfaces

import (
	"context"

	"github.com/abhinandkakkadi/rampnow-payment-svc/pkg/domain"
)

type PaymentUseCase interface {
	GetUserWalletById(ctx context.Context, userid int) (domain.UserWallet, error)
	CreateTransaction(ctx context.Context, transaction domain.Transaction) (domain.Transaction, error)
	GetAllTransactions(ctx context.Context) ([]domain.Transaction, error)
}
