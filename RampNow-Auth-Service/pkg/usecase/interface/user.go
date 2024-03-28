package interfaces

import (
	"context"

	domain "github.com/abhinandkakkadi/rampnow-auth-service/pkg/domain"
)

type UserUseCase interface {
	FindAll(ctx context.Context) ([]domain.Users, error)
	FindByID(ctx context.Context, id uint) (domain.Users, error)
	FindByName(ctx context.Context, email string) (domain.Users, error)
	Register(ctx context.Context, user domain.Users) (domain.Users, error)
	CreateWallet(ctx context.Context, userId int64) error
	Delete(ctx context.Context, id int64) error
	VerifyUser(ctx context.Context, email string, password string) error
}
