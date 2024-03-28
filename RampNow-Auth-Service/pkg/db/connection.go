package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	config "github.com/abhinandkakkadi/rampnow-auth-service/pkg/config"
	domain "github.com/abhinandkakkadi/rampnow-auth-service/pkg/domain"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	db.AutoMigrate(&domain.Users{})
	db.AutoMigrate(&domain.Transaction{})
	db.AutoMigrate(&domain.Wallet{})

	return db, dbErr
}
