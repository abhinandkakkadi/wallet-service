package payment

import (
	"github.com/abhinandkakkadi/rampnow/pkg/auth"
	"github.com/abhinandkakkadi/rampnow/pkg/config"
	"github.com/abhinandkakkadi/rampnow/pkg/payment/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) *ServiceClient {
	auth := auth.InitAuthMiddleware(authSvc)
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	order := r.Group("payment")

	order.Use(auth.AuthRequired)
	order.POST("/", svc.CreateTransaction)
	order.GET("/transactions", svc.GetAllTransactions)
	order.GET("/wallet_balance/:id", svc.GetWalletBalance)

	return svc
}

func (svc *ServiceClient) CreateTransaction(ctx *gin.Context) {
	routes.CreateTransaction(ctx, svc.Client)
}

func (svc *ServiceClient) GetAllTransactions(ctx *gin.Context) {
	routes.GetAllTransactions(ctx, svc.Client)
}

func (svc *ServiceClient) GetWalletBalance(ctx *gin.Context) {
	routes.GetWalletBalance(ctx, svc.Client)
}
