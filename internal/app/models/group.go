package models

import "time"

type Group struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(255);not null"`
	OwnerID   uint   `gorm:"not null"`
	Owner     User   `gorm:"foreignKey:OwnerID"`
	CreatedAt time.Time
}

type GroupMember struct {
	GroupID  uint      `gorm:"primaryKey"`
	UserID   uint      `gorm:"primaryKey"`
	Group    Group     `gorm:"foreignKey:GroupID"`
	User     User      `gorm:"foreignKey:UserID"`
	Role     string    `gorm:"type:varchar(255)"`
	JoinedAt time.Time `gorm:"not null"`
}
