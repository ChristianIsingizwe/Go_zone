package models

import "time"

type Order struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null;index"`
	Status    string `gorm:"not null;default:'pending'"`
	Total float64 `gorm:"not null;check:total >=0"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID"`
}


type OrderItem struct {
	ID uint `gorm:"primaryKey"`
	OrderID uint `gorm:"not null;index"`
	ProductID uint `gorm:"not null;index"`
	Quantity int `gorm:"not null;check:quantity > 0"`
	Price float64 `gorm:"not null;check:price >= 0"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`


	Order Order `gorm:"foreignKey:OrderID;constraints:OnDelete:CASCADE"`
	Product Product `gorm:"foreignKey:ProductID;constraints:OnDelete:CASCADE"`

}