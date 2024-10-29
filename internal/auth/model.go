package auth

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"unique;not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Role      string `gorm:"not null;default: 'customer'"`
	CreatedAt time.Time `gorm:"autoCreatedTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	ShoppingCart []CartItem `gorm:"foreignKey:UserID"`
	Orders []Order `gorm:"foreignKey:UserID"`
	Reviews []Review `gorm:"foreignKey:UserID"`
	Sessions []Session `gorm:"foreignKey:UserID"`
}