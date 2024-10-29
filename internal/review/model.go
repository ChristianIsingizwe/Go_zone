package review

type Review struct {
    ID        uint      `gorm:"primaryKey"`
    UserID    uint      `gorm:"not null;index"`
    ProductID uint      `gorm:"not null;index"`
    Rating    int       `gorm:"not null;check:rating BETWEEN 1 AND 5"`
    Comment   string    `gorm:"type:text"`
    CreatedAt time.Time `gorm:"autoCreateTime"`
    UpdatedAt time.Time `gorm:"autoUpdateTime"`

    User    User    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
    Product Product `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`
}