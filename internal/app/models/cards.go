package models

import "time"

// CardsExpense представляет таблицу cards_expenses.
type CardsExpense struct {
	ID          uint  `gorm:"primaryKey"`
	GroupID     uint  `gorm:"not null"`
	Group       Group `gorm:"foreignKey:GroupID"`
	PaidAt      time.Time
	Description string  `gorm:"type:text"`
	TotalAmount float64 `gorm:"type:numeric"`
	Currency    string  `gorm:"type:varchar(10)"`
	Settled     bool    `gorm:"default:false"`
	CreatedAt   time.Time
}

// CardsExpensePayer представляет таблицу cards_expenses_payers.
type CardsExpensePayer struct {
	ID             uint         `gorm:"primaryKey"`
	CardsExpenseID uint         `gorm:"not null"`
	CardsExpense   CardsExpense `gorm:"foreignKey:CardsExpenseID"`
	UserID         uint         `gorm:"not null"`
	User           User         `gorm:"foreignKey:UserID"`
	PaidAmount     float64      `gorm:"type:numeric"`
	PaidAt         time.Time
}

// CardsExpenseUser представляет таблицу cards_expenses_users.
type CardsExpenseUser struct {
	ID             uint         `gorm:"primaryKey"`
	CardsExpenseID uint         `gorm:"not null"`
	CardsExpense   CardsExpense `gorm:"foreignKey:CardsExpenseID"`
	UserID         uint         `gorm:"not null"`
	User           User         `gorm:"foreignKey:UserID"`
	ShareAmount    float64      `gorm:"type:numeric"`
	PaidAmount     float64      `gorm:"type:numeric"`
	PaidAt         time.Time
	CreatedAt      time.Time
}
