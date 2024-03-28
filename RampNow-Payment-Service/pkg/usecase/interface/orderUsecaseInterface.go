package interfaces

import (
	"context"

	"github.com/abhinandkakkadi/rampnow-payment-svc/pkg/domain"
	utils "github.com/abhinandkakkadi/rampnow-payment-svc/pkg/utils"
)

type OrderUseCase interface {
	CreateOrder(ctx context.Context, order domain.ReqOrder) (string, error)
	GetUserWalletById(ctx context.Context, userid int) (domain.UserWallet, error)
	CreateTransaction(ctx context.Context, transaction domain.Transaction) (domain.Transaction, error)
	GetAllTransactions(ctx context.Context) ([]domain.Transaction, error)
	UpdateOrder(ctx context.Context, orderid, status string) (string, error)
	FetchOrder(ctx context.Context, userid int, filter domain.Filter, pagenation utils.Filter) ([]domain.ReqOrder, utils.Metadata, error)
}
