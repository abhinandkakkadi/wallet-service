package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/abhinandkakkadi/rampnow-payment-svc/pkg/domain"
	repository "github.com/abhinandkakkadi/rampnow-payment-svc/pkg/repository/interface"
	interfaces "github.com/abhinandkakkadi/rampnow-payment-svc/pkg/usecase/interface"
	"github.com/google/uuid"
)

type paymentUseCase struct {
	orderRepo repository.PaymentRepository
}

func (o *paymentUseCase) GetAllTransactions(ctx context.Context) ([]domain.Transaction, error) {
	transactions, err := o.orderRepo.GetAllTransactions(ctx)
	return transactions, err
}

func (o *paymentUseCase) GetUserWalletById(ctx context.Context, userid int) (domain.UserWallet, error) {
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

func (o *paymentUseCase) CreateTransaction(ctx context.Context, transaction domain.Transaction) (domain.Transaction, error) {

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


func NewPaymentUseCase(repo repository.PaymentRepository) interfaces.PaymentUseCase {
	return &paymentUseCase{
		orderRepo: repo,
	}
}
