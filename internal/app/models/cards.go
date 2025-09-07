package models

import "time"

type CardsExpense struct {
	ID uint `gorm:"primaryKey"`

	GroupID     uint      `json:"group_id" gorm:"not null"`
	Group       Group     `json:"-" gorm:"foreignKey:GroupID"`
	PaidAt      time.Time `json:"paid_at"`
	Description string    `json:"description" gorm:"type:text"`
	TotalAmount float64   `json:"total_amount" gorm:"type:numeric"`
	Currency    string    `json:"currency" gorm:"type:varchar(10)"`
	Settled     bool      `json:"settled" gorm:"default:false"`

	CardsExpensePayers []CardsExpensePayer `gorm:"foreignKey:CardsExpenseID"`
	CardsExpenseUsers  []CardsExpenseUser  `gorm:"foreignKey:CardsExpenseID"`
	CreatedAt          time.Time           `gorm:"autoCreateTime"`
}

type CardsExpensePayer struct {
	ID uint `gorm:"primaryKey"`

	CardsExpenseID uint         `gorm:"not null"`
	CardsExpense   CardsExpense `json:"-" gorm:"foreignKey:CardsExpenseID"`
	UserID         uint         `gorm:"not null"`
	User           User         `json:"-" gorm:"foreignKey:UserID"`
	PaidAmount     float64      `json:"paid_amount" gorm:"type:numeric"`
	PaidAt         time.Time    `json:"paid_at"`
}

type CardsExpenseUser struct {
	ID uint `gorm:"primaryKey"`

	CardsExpenseID uint         `json:"cards_expense_id" gorm:"not null"`
	CardsExpense   CardsExpense `json:"-" gorm:"foreignKey:CardsExpenseID"`
	UserID         uint         `json:"user_id" gorm:"not null"`
	User           User         `json:"-" gorm:"foreignKey:UserID"`
	ShareAmount    float64      `json:"share_amount" gorm:"type:numeric"`
	PaidAmount     float64      `json:"paid_amount" gorm:"type:numeric"`
	PaidAt         time.Time    `json:"paid_at" gorm:"autoCreateTime"`
}
