package models

import (
	"gorm.io/gorm"
)

type GroupMember struct {
	gorm.Model
	ID      uint `gorm:"primaryKey"`
	GroupID uint `gorm:"column:group_id"`
	UserID  uint `gorm:"column:user_id"`

	Group Group `gorm:"foreignKey:GroupID"`
	User  User  `gorm:"foreignKey:UserID"`
}
