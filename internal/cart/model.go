package cart

import "time"

type CartItem struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"not null;index"`
	ProductID uint `gorm:"not null;index"`
	Quantity  int  `gorm:"not null;check:quantity > 0"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}