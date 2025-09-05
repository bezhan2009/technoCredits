package models

import "time"

type CardsExpense struct {
	ID uint `gorm:"primaryKey"`

	GroupID     uint  `gorm:"not null"`
	Group       Group `gorm:"foreignKey:GroupID"`
	PaidAt      time.Time
	Description string  `gorm:"type:text"`
	TotalAmount float64 `gorm:"type:numeric"`
	Currency    string  `gorm:"type:varchar(10)"`
	Settled     bool    `gorm:"default:false"`

	CardsExpensePayers []CardsExpensePayer `gorm:"foreignKey:CardsExpenseID"`
	CardsExpenseUsers  []CardsExpenseUser  `gorm:"foreignKey:CardsExpenseID"`
	CreatedAt          time.Time           `gorm:"autoCreateTime"`
}

type CardsExpensePayer struct {
	ID uint `gorm:"primaryKey"`

	CardsExpenseID uint         `gorm:"not null"`
	CardsExpense   CardsExpense `json:"-" gorm:"foreignKey:CardsExpenseID"`
	UserID         uint         `gorm:"not null"`
	User           User         `gorm:"foreignKey:UserID"`
	PaidAmount     float64      `gorm:"type:numeric"`
	PaidAt         time.Time
}

type CardsExpenseUser struct {
	ID uint `gorm:"primaryKey"`

	CardsExpenseID uint         `json:"cards_expense_id" gorm:"not null"`
	CardsExpense   CardsExpense `json:"-" gorm:"foreignKey:CardsExpenseID"`
	UserID         uint         `gorm:"not null"`
	User           User         `gorm:"foreignKey:UserID"`
	ShareAmount    float64      `gorm:"type:numeric"`
	PaidAmount     float64      `gorm:"type:numeric"`
	PaidAt         time.Time
}
