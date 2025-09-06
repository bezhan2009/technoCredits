package models

import "time"

type Group struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `json:"name" gorm:"type:varchar(255);not null"`
	OwnerID   uint   `json:"owner_id" gorm:"not null"`
	Owner     User   `json:"-" gorm:"foreignKey:OwnerID"`
	CreatedAt time.Time

	GroupMembers []GroupMember `gorm:"foreignKey:GroupID"`
}

type GroupMember struct {
	ID      uint  `gorm:"primaryKey"`
	GroupID uint  `json:"group_id" gorm:"primaryKey;autoIncrement:false"`
	UserID  uint  `json:"user_id" gorm:"primaryKey;autoIncrement:false"`
	Group   Group `json:"-" gorm:"foreignKey:GroupID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User    User  `json:"-" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	RoleID uint `json:"role_id" gorm:"not null"`
	Role   Role `json:"-" gorm:"foreignKey:RoleID"`

	JoinedAt time.Time `gorm:"not null"`
}
