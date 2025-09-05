package models

import (
	"gorm.io/gorm"
)

// CardsExpense - модель для таблицы "cards_expenses".
type CardsExpense struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	GroupID     uint   `gorm:"column:group_id"`
	PayedAt     int    `gorm:"column:payed_at"`
	Description string `gorm:"column:description"`
	Amount      int    `gorm:"column:amount"`

	// Связь с Group.
	Group Group `gorm:"foreignKey:GroupID;references:ID"`

	// One-to-many связь с CardsExpenseUser.
	CardsExpenseUsers []CardsExpenseUser `gorm:"foreignKey:CardsExpensesID;references:ID"`
}
