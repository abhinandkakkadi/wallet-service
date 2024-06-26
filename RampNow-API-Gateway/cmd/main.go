package main

import (
	"log"

	"github.com/abhinandkakkadi/rampnow/pkg/auth"
	"github.com/abhinandkakkadi/rampnow/pkg/config"
	"github.com/abhinandkakkadi/rampnow/pkg/payment"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title RampNow API
// @version 1.0
// @description This is RampNow API gateway for a wallet system. You can visit the GitHub repository at https://github.com/abhinandkakkadi/wallet-service

// @SecurityDefinition.Bearer BearerAuth
// @TokenUrl /auth/token
// @securityDefinitions.Bearer		type apiKey
// @securityDefinitions.Bearer		name Authorization
// @securityDefinitions.Bearer		in header
// @securityDefinitions.BasicAuth	type basic
func main() {
	// Load configuration
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	// Start router
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Register authentication and payment routes
	authSvc := *auth.RegisterRoutes(r, &c)
	payment.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
