package services

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/abhinandkakkadi/rampnow-payment-svc/pkg/domain"
	pb "github.com/abhinandkakkadi/rampnow-payment-svc/pkg/pb"
	usecase "github.com/abhinandkakkadi/rampnow-payment-svc/pkg/usecase/interface"
)

type PaymentService struct {
	orderUseCase usecase.PaymentUseCase
}

func (c *PaymentService) GetWalletBalance(ctx context.Context, req *pb.GetWalletBalanceRequest) (*pb.GetWalletBalanceResponse, error) {
	// Get 
	if req.Id == 0 {
		return &pb.GetWalletBalanceResponse{
			Status: http.StatusBadRequest,
			Error:  "Invalid ID",
		}, nil
	}

	userWallet, err := c.orderUseCase.GetUserWalletById(ctx, int(req.Id))
	if err != nil {
		return &pb.GetWalletBalanceResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  fmt.Sprint(errors.New("unable to fetch data")),
		}, errors.New(err.Error())
	}

	data := &pb.UserWallet{
		FullName: userWallet.FullName,
		RampId: userWallet.RampId,
		WalletBalance: fmt.Sprint(userWallet.WalletBalance),
	}

	fmt.Printf("wallet balance: %+v ", data)

	return &pb.GetWalletBalanceResponse{
		Status: http.StatusOK,
		Data: data,
	}, nil


} 

func (c *PaymentService) GetTransactions(ctx context.Context, req *pb.GetTransactionRequest) (*pb.GetTransactionResponse, error) {
	transactions, err := c.orderUseCase.GetAllTransactions(ctx)
	if err != nil {
		return &pb.GetTransactionResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  fmt.Sprint(errors.New("unable to fetch data")),
		}, errors.New(err.Error())
	}

	var trans []*pb.Transaction
	for _, transaction := range transactions {
		trans1 := &pb.Transaction{
			Id:        transaction.Id,
			OrderId: transaction.PayeeRampId,
			PayerRampId: transaction.PayerRampId,
			PayeeRampId: transaction.PayeeRampId,
			PaymentAmount: fmt.Sprint(transaction.PaymentAmount),
		}
		trans = append(trans, trans1)
	}

	// Return the response
	return &pb.GetTransactionResponse{
		Status: http.StatusOK,
		Transaction:   trans,
	}, nil
}

func (c *PaymentService) CreateTransaction(ctx context.Context, req *pb.CreateTransactionRequest) (*pb.CreateTransactionResponse, error) {
	// Create transaction 
	transaction := domain.Transaction{
		PayerRampId: req.PayerRampId,
		PayeeRampId: req.PayeeRampId,
		PaymentAmount: req.PaymentAmount,
	}

	transaction, err := c.orderUseCase.CreateTransaction(ctx, transaction)
	if err != nil {
		return &pb.CreateTransactionResponse{
			Status: http.StatusUnprocessableEntity,
			Error: err.Error(),
		}, nil
	}

	return &pb.CreateTransactionResponse{
		Status: http.StatusCreated,
		Id: transaction.OrderId,
	}, nil
}




func NewPaymentService(usecase usecase.PaymentUseCase) *PaymentService {
	return &PaymentService{
		orderUseCase: usecase,
	}
}
