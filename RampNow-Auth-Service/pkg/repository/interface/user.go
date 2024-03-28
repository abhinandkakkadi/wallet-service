package interfaces

import (
	"context"

	"github.com/abhinandkakkadi/rampnow-auth-service/pkg/domain"
)

type UserRepository interface {
	FindAll(ctx context.Context) ([]domain.Users, error)
	FindByID(ctx context.Context, id uint) (domain.Users, error)
	CreateWallet(ctx context.Context, wallet domain.Wallet) (error)
	FindByName(ctx context.Context, email string) (domain.Users, error)
	Save(ctx context.Context, user domain.Users) (domain.Users, error)
	Delete(ctx context.Context, id int64) error
	FindPassword(ctx context.Context, email string) (string, error)
}
