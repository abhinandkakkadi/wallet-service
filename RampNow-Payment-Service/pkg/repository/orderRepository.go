package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/abhinandkakkadi/rampnow-payment-svc/pkg/domain"
	interfaces "github.com/abhinandkakkadi/rampnow-payment-svc/pkg/repository/interface"
	utils "github.com/abhinandkakkadi/rampnow-payment-svc/pkg/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type orderDatabase struct {
	DB *gorm.DB
}

func (o *orderDatabase) GetUserIdFromRampId(ctx context.Context, rampId string) (int, error) {
	var userID int 

	err := o.DB.Raw("SELECT id FROM users WHERE ramp_id = ?", rampId).Scan(&userID).Error
	if err != nil {
		return 0, errors.New("transaction failed due to server inconsistencies")
	}

	return userID, nil
}

func (o *orderDatabase) GetWalletAmountFromUserId(ctx context.Context, userId int) (float64, error) {
	var walletAmount float64 

	err := o.DB.Raw("SELECT wallet_balance FROM wallets WHERE  user_id = ?", userId).Scan(&walletAmount).Error
	if err != nil {
		return 0, errors.New("transaction failed due to server inconsistencies")
	}

	return walletAmount, nil
}

func (o *orderDatabase) IncrementWalletBalance(ctx context.Context, userId int, paymentAmount float64) (error) {
	err := o.DB.Exec("UPDATE wallets SET wallet_balance = wallet_balance + ? WHERE user_id = ?", paymentAmount, userId).Error
	return err
}

func (o *orderDatabase) DecrementWalletBalance(ctx context.Context, userId int, paymentAmount float64) (error) {
	err := o.DB.Exec("UPDATE wallets SET wallet_balance = wallet_balance - ? WHERE user_id = ?", paymentAmount, userId).Error
	return err
}

func (o *orderDatabase) FindUser(ctx context.Context, rampId string) (int, error) {
	var exists int

	err := o.DB.Raw("SELECT COUNT(*) FROM users WHERE ramp_id = ?", rampId).Scan(&exists).Error
	if err != nil {
		return 0, errors.New("transaction failed due to server inconsistencies")
	}

	return exists, nil
}

func (o *orderDatabase) FindByID(ctx context.Context, userId int) (domain.Users, error) {
	var user domain.Users
	err := o.DB.First(&user, userId).Error

	return user, err
}

func (o *orderDatabase) GetAllTransactions(ctx context.Context) ([]domain.Transaction, error) {
	var transaction []domain.Transaction
	err := o.DB.Find(&transaction).Error

	return transaction, err
}



func (o *orderDatabase) CreateTransaction(ctx context.Context, transaction domain.Transaction) (domain.Transaction, error) {
	err := o.DB.Save(&transaction).Error
	return transaction, err
}

// FetchOrder implements interfaces.OrderRepository
func (o *orderDatabase) FetchOrder(ctx context.Context, userid int, filter domain.Filter, pagenation utils.Filter) ([]domain.Order, utils.Metadata, error) {

	sql := "SELECT * FROM orders WHERE 1=1"
	if filter.Status != "" {
		sql += fmt.Sprintf(" AND status='%s'", filter.Status)
	}
	if filter.MinTotal != 0 {
		sql += fmt.Sprintf(" AND total>=%v", filter.MinTotal)
	}
	if filter.MaxTotal != 0 {
		sql += fmt.Sprintf(" AND total<=%v", filter.MaxTotal)
	}
	if filter.SortBy != "" {
		sql += fmt.Sprintf(" ORDER BY %s", filter.SortBy)
		if filter.SortOrder == "desc" {
			sql += " DESC"
		}
	}

	sql += fmt.Sprintf(" LIMIT %v OFFSET %v", pagenation.Limit(), pagenation.Offset())

	order := []domain.Order{}
	var TotalRecords int64
	err := o.DB.Raw(sql).Scan(&order).Count(&TotalRecords).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("no order in the list")
	}

	return order, utils.ComputeMetaData(int(TotalRecords), pagenation.Page, pagenation.PageSize), err
}

// UpdateOrder implements interfaces.OrderRepository
func (o *orderDatabase) UpdateOrder(ctx context.Context, orderid string, status string) (string, error) {
	order := domain.Order{}
	err := o.DB.Model(&order).Where("id = ?", orderid).Update("status", status).Error
	return orderid, err
}

// FindItem implements interfaces.OrderRepository
func (c *orderDatabase) FindItem(ctx context.Context, id string) (domain.Item, error) {

	var item domain.Item
	err := c.DB.Where("id = ?", id).First(&item).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("no item in the list")
	}

	return item, err
}

// CreateItem implements interfaces.OrderRepository
func (o *orderDatabase) CreateItem(ctx context.Context, item domain.Item) (string, error) {
	err := o.DB.Create(&item).Error
	return item.ID, err
}

// CreateOrder implements interfaces.OrderRepository
func (o *orderDatabase) CreateOrder(ctx context.Context, order domain.Order) (string, error) {
	order.ID = uuid.New().String()
	err := o.DB.Create(&order).Error
	return order.ID, err
}

func NewOrderRepository(DB *gorm.DB) interfaces.OrderRepository {
	return &orderDatabase{DB}
}
