package auth

import (
	"github.com/abhinandkakkadi/rampnow/pkg/auth/routes"
	"github.com/abhinandkakkadi/rampnow/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}
	authMiddleware := InitAuthMiddleware(svc)

	// Create a new group for the auth routes that require authentication
	authRoutes := r.Group("/auth")
	authRoutes.Use(authMiddleware.RefreshTokenMiddleware)
	authRoutes.POST("/token-refresh", svc.TokenRefresh)

	routes := r.Group("/auth")
	routes.POST("/register", svc.Register)
	routes.POST("/login", svc.Login)

	userRoutes := r.Group("/user")
	userRoutes.GET("/finduser/:id", svc.FindUser)
	userRoutes.GET("/getusers", svc.GetUsers)
	return svc
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}


func (svc *ServiceClient) GetUsers(ctx *gin.Context) {
	routes.GetUsers(ctx, svc.Client)
}
func (svc *ServiceClient) FindUser(ctx *gin.Context) {
	routes.FindUser(ctx, svc.Client)
}
func (svc *ServiceClient) TokenRefresh(ctx *gin.Context) {
	routes.TokenRefresh(ctx, svc.Client)
}
