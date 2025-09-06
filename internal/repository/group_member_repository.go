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
	return TranslateGormError(r.DB.Create(member).Error)
}

func (r *GroupMemberRepository) GetByGroupID(groupID uint) ([]models.GroupMember, error) {
	var members []models.GroupMember
	err := r.DB.Where("group_id = ?", groupID).Find(&members).Error
	if err != nil {
		return members, TranslateGormError(err)
	}

	return members, nil
}

func (r *GroupMemberRepository) UpdateRole(groupID uint, userID uint, role uint) error {
	return TranslateGormError(r.DB.Model(&models.GroupMember{}).
		Where("group_id = ? AND user_id = ?", groupID, userID).
		Update("role_id", role).Error)
}

func (r *GroupMemberRepository) Delete(groupID uint, userID uint) error {
	return TranslateGormError(r.DB.Where("group_id = ? AND user_id = ?", groupID, userID).
		Delete(&models.GroupMember{}).Error)
}
