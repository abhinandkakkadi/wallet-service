package main

import (
	"log"

	config "github.com/abhinandkakkadi/rampnow-auth-service/pkg/config"
	di "github.com/abhinandkakkadi/rampnow-auth-service/pkg/di"
)

func main() {
	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}

	diErr := di.InitializeAPI(config)
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} 
}
