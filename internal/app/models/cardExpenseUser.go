package models

import (
	"gorm.io/gorm"
)

type CardsExpenseUser struct {
	gorm.Model
	ID              uint `gorm:"primaryKey"`
	UserID          uint `gorm:"column:user_id"`
	Cost            int  `gorm:"column:cost"`
	Spend           int  `gorm:"column:spend"`
	PayedAt         int  `gorm:"column:payed_at"`
	CardsExpensesID uint `gorm:"column:cards_expenses_id"`

	User User `gorm:"foreignKey:UserID;references:ID"`

	CardsExpense CardsExpense `gorm:"foreignKey:CardsExpensesID;references:ID"`
}
