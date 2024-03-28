package main

import (
	"fmt"
	"log"

	"github.com/abhinandkakkadi/rampnow-payment-svc/pkg/config"
	"github.com/abhinandkakkadi/rampnow-payment-svc/pkg/di"
)

func main() {
	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}
	fmt.Println("config : ", config)

	server, diErr := di.InitializeAPI(config)
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()

	}

}
