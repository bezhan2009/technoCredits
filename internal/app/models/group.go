package models

import (
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	ID      uint   `gorm:"primaryKey"`
	Name    string `gorm:"column:name"`
	OwnerID uint   `gorm:"column:owner_id"`

	Owner User `gorm:"foreignKey:OwnerID;references:ID"`

	Users []User `gorm:"many2many:groups_member;foreignKey:ID;references:ID"`

	CardsExpenses []CardsExpense `gorm:"foreignKey:GroupID;references:ID"`
}
