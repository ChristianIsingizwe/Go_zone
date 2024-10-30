package services

import (
	"fmt"
	"log"
	"os"

	"github.com/ChristianIsingizwe/Go_zone/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() error{
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database : %v", err)
	}

	if err := database.AutoMigrate(&models.User{}, &models.Session{}, &models.CartItem{}, &models.Order{}, &models.OrderItem{}, &models.Review{}); err != nil {
		log.Fatalf("Failed to auto-migrate:  %v", err)
	}
	DB = database

	return nil 
}
