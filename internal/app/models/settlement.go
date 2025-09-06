package models

import "time"

type Settlement struct {
	ID         uint    `gorm:"primaryKey"`
	GroupID    uint    `gorm:"not null"`
	Group      Group   `json:"-" gorm:"foreignKey:GroupID"`
	FromUserID uint    `gorm:"not null"`
	FromUser   User    `json:"from_user" gorm:"foreignKey:FromUserID"`
	ToUserID   uint    `json:"to_user_id" gorm:"not null"`
	ToUser     User    `json:"-" gorm:"foreignKey:ToUserID"`
	Amount     float64 `json:"amount" gorm:"type:numeric"`
	Currency   string  `json:"currency" gorm:"type:varchar(10);default:TJS"`
	CreatedAt  time.Time
	Note       string `json:"note" gorm:"type:text"`
}
