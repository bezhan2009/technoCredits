package repository

import (
	"gorm.io/gorm"
	"technoCredits/internal/app/models"
)

type GroupMemberRepository struct {
	DB *gorm.DB
}

func NewGroupMemberRepository(db *gorm.DB) *GroupMemberRepository {
	return &GroupMemberRepository{DB: db}
}

func (r *GroupMemberRepository) Create(member *models.GroupMember) error {
	return r.DB.Create(member).Error
}

func (r *GroupMemberRepository) GetByGroupID(groupID uint) ([]models.GroupMember, error) {
	var members []models.GroupMember
	err := r.DB.Where("group_id = ?", groupID).Find(&members).Error
	return members, err
}

func (r *GroupMemberRepository) UpdateRole(groupID uint, userID uint, role string) error {
	return r.DB.Model(&models.GroupMember{}).
		Where("group_id = ? AND user_id = ?", groupID, userID).
		Update("role", role).Error
}

func (r *GroupMemberRepository) Delete(groupID uint, userID uint) error {
	return r.DB.Where("group_id = ? AND user_id = ?", groupID, userID).
		Delete(&models.GroupMember{}).Error
}
