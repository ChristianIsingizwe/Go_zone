package sessions

import "time"

type Session struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null;index"`
	JwtToken  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}