//go:build wireinject
// +build wireinject

package di

import (
	"github.com/SethukumarJ/sellerapp-order-svc/pkg/api/services"
	"github.com/SethukumarJ/sellerapp-order-svc/pkg/config"
	"github.com/SethukumarJ/sellerapp-order-svc/pkg/db"
	"github.com/SethukumarJ/sellerapp-order-svc/pkg/repository"
	"github.com/SethukumarJ/sellerapp-order-svc/pkg/usecase"

	http "github.com/SethukumarJ/sellerapp-order-svc/pkg/api"
	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase,
		repository.NewOrderRepository,
		usecase.NewOrderUseCase,
		services.NewOrderService,
		http.NewServerHTTP,
	)

	return &http.ServerHTTP{}, nil
}
