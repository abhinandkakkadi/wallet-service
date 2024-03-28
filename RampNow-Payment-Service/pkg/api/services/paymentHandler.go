package services

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/abhinandkakkadi/rampnow-payment-svc/pkg/domain"
	pb "github.com/abhinandkakkadi/rampnow-payment-svc/pkg/pb"
	usecase "github.com/abhinandkakkadi/rampnow-payment-svc/pkg/usecase/interface"
	utils "github.com/abhinandkakkadi/rampnow-payment-svc/pkg/utils"
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



// UpdateOrder implements pb.OrderServiceServer
func (c *PaymentService) UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error) {
	id, err := c.orderUseCase.UpdateOrder(ctx, req.OrderId, req.Status)
	if err != nil {
		return &pb.UpdateOrderResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  err.Error(),
		}, err
	}

	return &pb.UpdateOrderResponse{
		Status: http.StatusOK,
		Id:     id,
	}, nil

}

// FetchOrder implements pb.OrderServiceServer
func (c *PaymentService) FetchOrder(ctx context.Context, req *pb.FetchOrderRequest) (*pb.FetchOrderResponse, error) {
	filter := domain.Filter{
		Status:    req.Status,
		MinTotal:  float64(req.MinTotal),
		MaxTotal:  float64(req.MaxTotal),
		SortBy:    req.SortBy,
		SortOrder: req.SortOrder,
	}

	pagnation := utils.Filter{
		Page:     int(req.Filter.Page),
		PageSize: int(req.Filter.PageSize),
	}

	orders, Metadata, err := c.orderUseCase.FetchOrder(ctx, int(req.UserId), filter, pagnation)
	if err != nil {
		return &pb.FetchOrderResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  err.Error(),
		}, err
	}
	var Od []*pb.Order
	for _, order := range orders {
		var reitems []*pb.Item
		for _, items := range order.Item {
			items := pb.Item{
				ID:          items.ID,
				Description: items.Description,
				Price:       float32(items.Price),
				Quantity:    int64(items.Quantity),
			}
			reitems = append(reitems, &items)
		}

		od := pb.Order{
			OrderId:      order.ID,
			Status:       order.Status,
			Item:         reitems,
			Total:        float32(order.Total),
			CurrencyUnit: order.CurrencyUnit,
		}
		Od = append(Od, &od)
	}

	return &pb.FetchOrderResponse{
		Status: http.StatusOK,
		Orders: Od,
		Metadata: &pb.Metadata{
			CurrentPage:  int64(Metadata.CurrentPage),
			PageSize:     int64(Metadata.PageSize),
			FirstPage:    int64(Metadata.FirstPage),
			LastPage:     int64(Metadata.LastPage),
			TotalRecords: int64(Metadata.TotalRecords),
		},
	}, nil

}

func NewPaymentService(usecase usecase.PaymentUseCase) *PaymentService {
	return &PaymentService{
		orderUseCase: usecase,
	}
}
