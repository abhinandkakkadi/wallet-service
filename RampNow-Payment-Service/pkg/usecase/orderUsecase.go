package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/abhinandkakkadi/rampnow-payment-svc/pkg/domain"
	repository "github.com/abhinandkakkadi/rampnow-payment-svc/pkg/repository/interface"
	interfaces "github.com/abhinandkakkadi/rampnow-payment-svc/pkg/usecase/interface"
	utils "github.com/abhinandkakkadi/rampnow-payment-svc/pkg/utils"
	"github.com/google/uuid"
)

type orderUseCase struct {
	orderRepo repository.OrderRepository
}

func (o *orderUseCase) GetAllTransactions(ctx context.Context) ([]domain.Transaction, error) {
	transactions, err := o.orderRepo.GetAllTransactions(ctx)
	return transactions, err
}

func (o *orderUseCase) GetUserWalletById(ctx context.Context, userid int) (domain.UserWallet, error) {
	// Get name and ramp id 
	user, err := o.orderRepo.FindByID(ctx, userid)
	if err != nil {
		return domain.UserWallet{}, err
	}

	// Get wallet balance of this user
	// Get WalletAmountByUserId
	walletAmount, err := o.orderRepo.GetWalletAmountFromUserId(ctx, userid)
	if err != nil {
		return domain.UserWallet{}, err
	}
	
	userWallet := domain.UserWallet{
		FullName: user.FullName,
		RampId: user.RampId,
		WalletBalance: float32(walletAmount),
	}

	return userWallet, nil
}

func (o *orderUseCase) CreateTransaction(ctx context.Context, transaction domain.Transaction) (domain.Transaction, error) {

	transaction.PaymentStatus = "PAID"
	transaction.PaymentDate = time.Now()
	transaction.OrderId = uuid.New().String()

	// Check if amount is less than 0
	if transaction.PaymentAmount <= 0 {
		return domain.Transaction{}, errors.New("transaction amount should be greater than 0")
	}

	// Check if payer and payee are not same person
	if transaction.PayerRampId == transaction.PayeeRampId {
		return domain.Transaction{}, errors.New("payer and payee should not be the same person")
	}

	// Check if payer exists
	payerExists, err := o.orderRepo.FindUser(ctx, transaction.PayerRampId)
	if err != nil {
		return domain.Transaction{}, err
	}

	// Check if payee exists
	payeeExists, err := o.orderRepo.FindUser(ctx, transaction.PayeeRampId)
	if err != nil {
		return domain.Transaction{}, err
	}

	if payerExists == 0 {
		return domain.Transaction{}, errors.New("payer rampid invalid")
	}
	if payeeExists == 0 {
		return domain.Transaction{}, errors.New("payee rampid invalid")
	}

	// Check payer wallet balance
	payerUserId, err := o.orderRepo.GetUserIdFromRampId(ctx, transaction.PayerRampId)
	if err != nil {
		return domain.Transaction{}, err
	}

	payeeUserId, err := o.orderRepo.GetUserIdFromRampId(ctx, transaction.PayeeRampId)
	if err != nil {
		return domain.Transaction{}, err
	}

	// Get WalletAmountByUserId
	walletAmount, err := o.orderRepo.GetWalletAmountFromUserId(ctx, payerUserId)
	if err != nil {
		return domain.Transaction{}, err
	}

	// Check if the payer wallet have sufficient amount
	if walletAmount < float64(transaction.PaymentAmount) {
		return domain.Transaction{}, errors.New("insufficient fund in the wallet")
	}

	// Update safe of payee and payer
	err = o.orderRepo.IncrementWalletBalance(ctx, payeeUserId, float64(transaction.PaymentAmount))
	if err != nil {
		return domain.Transaction{}, err
	}

	// Update safe of payee and payer
	err = o.orderRepo.DecrementWalletBalance(ctx, payerUserId, float64(transaction.PaymentAmount))
	if err != nil {
		return domain.Transaction{}, err
	}

	transaction, err = o.orderRepo.CreateTransaction(ctx, transaction)
	return transaction, err
}

func (o *orderUseCase) FetchOrder(ctx context.Context, userid int, filter domain.Filter, pagenation utils.Filter) ([]domain.ReqOrder, utils.Metadata, error) {
	order, metadata, err := o.orderRepo.FetchOrder(ctx, userid, filter, pagenation)
	if err != nil {
		return []domain.ReqOrder{}, metadata, err
	}

	Rorder := []domain.ReqOrder{}
	for _, od := range order {
		fmt.Println("itemid", od.Item_id)
		itemId := strings.Split(od.Item_id, ",")
		items := []domain.Item{}
		for _, id := range itemId {
			item, err := o.orderRepo.FindItem(ctx, id)
			if err == nil {
				items = append(items, item)
			}

		}
		recorder := domain.ReqOrder{
			ID:           od.ID,
			Status:       od.Status,
			Item:         items,
			Total:        od.Total,
			CurrencyUnit: od.CurrencyUnit,
		}
		Rorder = append(Rorder, recorder)

	}

	return Rorder, metadata, err
}

// UpdateOrder implements interfaces.OrderUseCase
func (o *orderUseCase) UpdateOrder(ctx context.Context, orderid string, status string) (string, error) {
	id, err := o.orderRepo.UpdateOrder(ctx, orderid, status)
	return id, err
}

// CreateOrder implements interfaces.OrderUseCase
func (o *orderUseCase) CreateOrder(ctx context.Context, order domain.ReqOrder) (string, error) {
	items := order.Item
	fmt.Println(len(items))
	if len(items) == 0 {
		return "", errors.New("there is no item in order")
	}

	var it []string
	var id string

	for _, item := range items {
		_, err := o.orderRepo.FindItem(ctx, item.ID)
		if err != nil {
			id, err = o.orderRepo.CreateItem(ctx, item)
			if err != nil {
				return "", err
			}
		}

		it = append(it, id)
	}
	itemIDs := strings.Join(it, ",")
	od := domain.Order{
		ID:           order.ID,
		Status:       order.Status,
		Item_id:      itemIDs,
		Total:        order.Total,
		CurrencyUnit: order.CurrencyUnit,
	}
	id, err := o.orderRepo.CreateOrder(ctx, od)
	return id, err
}

func NewOrderUseCase(repo repository.OrderRepository) interfaces.OrderUseCase {
	return &orderUseCase{
		orderRepo: repo,
	}
}
