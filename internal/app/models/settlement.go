package models

import "time"

type Settlement struct {
	ID         uint    `gorm:"primaryKey"`
	GroupID    uint    `gorm:"not null"`
	Group      Group   `gorm:"foreignKey:GroupID"`
	FromUserID uint    `gorm:"not null"`
	FromUser   User    `gorm:"foreignKey:FromUserID"`
	ToUserID   uint    `gorm:"not null"`
	ToUser     User    `gorm:"foreignKey:ToUserID"`
	Amount     float64 `gorm:"type:numeric"`
	Currency   string  `gorm:"type:varchar(10)"`
	CreatedAt  time.Time
	Note       string `gorm:"type:text"`
}
