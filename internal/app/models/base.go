package models

import (
	"database/sql"
	"time"
)

type GormDeletedAt sql.NullTime

// GormModel содержит общие поля для всех таблиц.
type GormModel struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt GormDeletedAt `gorm:"index"`
}

// User представляет таблицу users.
type User struct {
	ID        uint   `gorm:"primaryKey"`
	RoleID    uint   `json:"role_id" gorm:"not null"`
	Role      Role   `json:"-" gorm:"foreignKey:RoleID"`
	FullName  string `json:"full_name" gorm:"type:varchar(255)"`
	Username  string `json:"username" gorm:"type:varchar(255)"`
	Email     string `json:"email" gorm:"type:varchar(255);unique;not null"`
	Password  string `json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
