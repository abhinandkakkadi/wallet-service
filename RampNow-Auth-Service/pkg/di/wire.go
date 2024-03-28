//go:build wireinject
// +build wireinject

package di

import (
	http "github.com/SethukumarJ/sellerapp-auth-service/pkg/api"
	handler "github.com/SethukumarJ/sellerapp-auth-service/pkg/api/handler"
	config "github.com/SethukumarJ/sellerapp-auth-service/pkg/config"
	db "github.com/SethukumarJ/sellerapp-auth-service/pkg/db"
	repository "github.com/SethukumarJ/sellerapp-auth-service/pkg/repository"
	usecase "github.com/SethukumarJ/sellerapp-auth-service/pkg/usecase"
	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase,
		repository.NewUserRepository,
		usecase.NewUserUseCase,
		usecase.NewJWTUsecase,
		handler.NewUserHandler,
		http.NewServerHTTP)

	return &http.ServerHTTP{}, nil
}
