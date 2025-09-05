package models

import (
	"time"

	"gorm.io/gorm"
)

// GormModel содержит общие поля для всех таблиц.
type GormModel struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// User представляет таблицу users.
type User struct {
	ID        uint   `gorm:"primaryKey"`
	RoleID    uint   `gorm:"not null"`
	Role      Role   `gorm:"foreignKey:RoleID"`
	FullName  string `gorm:"type:varchar(255)"`
	Username  string `json:"username" gorm:"type:varchar(255)"`
	Password  string `json:"password"`
	Email     string `gorm:"type:varchar(255);unique;not null"`
	Login     string `gorm:"type:varchar(255);unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
