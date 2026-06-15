package config

import (
	"fmt"
	"os"

	"github.com/payment-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

var DB *gorm.DB

func ConfigDatabase() {
	pgConfig := PostgresConfig{
		Host:     string(os.Getenv("DB_HOST")),
		Port:     string(os.Getenv("DB_PORT")),
		User:     string(os.Getenv("DB_USER")),
		Password: string(os.Getenv("DB_PASSWORD")),
		Dbname:   string(os.Getenv("DB_NAME")),
	}

	// Set postgres whitelist config
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=require",
		pgConfig.User,
		pgConfig.Password,
		pgConfig.Host,
		pgConfig.Port,
		pgConfig.Dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("failed to connect database")
	}

	err = db.AutoMigrate(
		&models.Order{},
		&models.Product{},
		&models.CheckoutRequest{},
		&models.User{},
	)

	if err != nil {
		panic(err)
	}

	DB = db
}
