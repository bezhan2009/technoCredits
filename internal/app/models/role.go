package models

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`

	// Один Role может иметь много User.
	Users []User `gorm:"foreignKey:RoleID;references:ID"`
}
