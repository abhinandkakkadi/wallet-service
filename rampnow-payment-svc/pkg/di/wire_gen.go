// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/abhinandkakkadi/rampnow-payment-svc/pkg/api"
	"github.com/abhinandkakkadi/rampnow-payment-svc/pkg/api/services"
	"github.com/abhinandkakkadi/rampnow-payment-svc/pkg/config"
	"github.com/abhinandkakkadi/rampnow-payment-svc/pkg/db"
	"github.com/abhinandkakkadi/rampnow-payment-svc/pkg/repository"
	"github.com/abhinandkakkadi/rampnow-payment-svc/pkg/usecase"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return err
	}
	paymentRepository := repository.NewPaymentRepository(gormDB)
	paymentUseCase := usecase.NewPaymentUseCase(paymentRepository)
	paymentService := services.NewPaymentService(paymentUseCase)
	http.NewServerHTTP(paymentService)
	return  nil
}
