package main

import (
	"fmt"
	"log"

	"github.com/abhinandkakkadi/rampnow/pkg/auth"
	"github.com/abhinandkakkadi/rampnow/pkg/config"
	order "github.com/abhinandkakkadi/rampnow/pkg/order_svc"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Order Management API
// @version 1.0
// @description This is order management sample service. You can visit the GitHub repository at https://github.com/abhinandkakkadi/rampnow-Gateway

// @contact.name API Support
// @contact.url sethukumarj.com
// @contact.email sethukumarj.76@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securityDefinitions.apikey BearerAuth
// @in header
// @name authorization

// @host localhost:3005
// @BasePath /
// @query.collection.format multi
func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	authSvc := *auth.RegisterRoutes(r, &c)
	fmt.Println("authSvc", authSvc)
	order.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
