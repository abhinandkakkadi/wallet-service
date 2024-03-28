package order

import (
	"github.com/abhinandkakkadi/rampnow/pkg/auth"
	"github.com/abhinandkakkadi/rampnow/pkg/config"
	"github.com/abhinandkakkadi/rampnow/pkg/order_svc/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) *ServiceClient {
	auth := auth.InitAuthMiddleware(authSvc)
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	order := r.Group("transaction")

	order.Use(auth.AuthRequired)
	order.POST("/", svc.CreateOrder)
	order.GET("/transactions", svc.UpdateOrder)
	order.GET("/wallet_balance/:id", svc.FetchOrder)

	return svc
}

func (svc *ServiceClient) CreateOrder(ctx *gin.Context) {
	routes.CreateOrder(ctx, svc.Client)
}

func (svc *ServiceClient) FetchOrder(ctx *gin.Context) {
	routes.FetchOrder(ctx, svc.Client)
}

func (svc *ServiceClient) UpdateOrder(ctx *gin.Context) {
	routes.UpdateOrder(ctx, svc.Client)
}
