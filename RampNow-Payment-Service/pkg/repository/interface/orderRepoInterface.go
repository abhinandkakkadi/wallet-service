package interfaces

import (
	"context"

	"github.com/abhinandkakkadi/rampnow-payment-svc/pkg/domain"
	utils "github.com/abhinandkakkadi/rampnow-payment-svc/pkg/utils"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order domain.Order) (string, error)
	FindByID(ctx context.Context, userId int) (domain.Users, error)
	GetAllTransactions(ctx context.Context) ([]domain.Transaction, error)
	CreateTransaction(ctx context.Context, order domain.Transaction) (domain.Transaction, error)
	FindUser(ctx context.Context, rampId string) (int, error)
	GetUserIdFromRampId(ctx context.Context, rampId string) (int, error)
	GetWalletAmountFromUserId(ctx context.Context, userId int) (float64, error)
	IncrementWalletBalance(ctx context.Context, userId int, paymentAmount float64) (error)
	DecrementWalletBalance(ctx context.Context, userId int, paymentAmount float64) (error)
	CreateItem(ctx context.Context, item domain.Item) (string, error)
	FindItem(ctx context.Context, id string) (domain.Item, error)
	UpdateOrder(ctx context.Context, orderid, status string) (string, error)
	FetchOrder(ctx context.Context, userid int, filter domain.Filter, pagenation utils.Filter) ([]domain.Order, utils.Metadata, error)
}
