package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           uint   `gorm:"primaryKey"`
	FullName     string `gorm:"column:full_name"`
	RoleID       uint   `gorm:"primaryKey;column:role_id"`
	PasswordHash string `gorm:"column:password_hash"`
	Email        string `gorm:"column:email"`
	Login        string `gorm:"column:login"`

	Role   Role
	Groups []Group `gorm:"many2many:groups_member;foreignKey:ID;references:ID"`
}
