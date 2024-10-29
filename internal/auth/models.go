package auth

import "time"

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Username  string `json:"username" gorm:"unique;not null"`
	Email     string `json:"email" gorm:"unique;not null"`
	Password  string `json:"-"`
	Role      string `json:"role" gorm:"not null;default:'customer'"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt time.Time `json:"updated_at"`
	ShoppingCart []CartItem `json:"shopping_cart"`
	Orders []Order `json:"orders"`
}