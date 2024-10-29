package models

import "time"

type Product struct {
	ID            uint    `gorm:"primaryKey"`
	Name          string  `gorm:"not null"`
	Description   string  `gorm:"not null"`
	Price         float64 `gorm:"not null;check:price>=0"`
	stockQuantity int     `gorm:"not null;check:stock_quantity >= 0"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	CartItems []CartItem `gorm:"foreignKey:ProductID"`
	OrderItems []OrderItem `gorm:"foreignKey:ProductID"`
	Reviews []Review `gorm:"foreignKey:ProductID"`
}