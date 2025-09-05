package models

import "time"

// Group представляет таблицу groups.
type Group struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(255);not null"`
	OwnerID   uint   `gorm:"not null"`
	Owner     User   `gorm:"foreignKey:OwnerID"`
	CreatedAt time.Time
}

// GroupMember представляет таблицу groups_member, которая является связующей таблицей для отношений многие ко многим.
type GroupMember struct {
	GroupID  uint      `gorm:"primaryKey;autoIncrement:false"`
	UserID   uint      `gorm:"primaryKey;autoIncrement:false"`
	Group    Group     `gorm:"foreignKey:GroupID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User     User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Role     string    `gorm:"type:varchar(255)"`
	JoinedAt time.Time `gorm:"not null"`
}
