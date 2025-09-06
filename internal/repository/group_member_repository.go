package repository

import (
	"gorm.io/gorm"
	"technoCredits/internal/app/models"
	"technoCredits/pkg/db"
	"technoCredits/pkg/logger"
)

func GetGroupMembersByGroupID(groupID uint) (groupMembers []models.GroupMember, err error) {
	if err = db.GetDBConn().Model(&models.GroupMember{}).Where("group_id = ?", groupID).Find(&groupMembers).Error; err != nil {
		logger.Error.Printf("[repository.GetGroupMembersByGroupID] Error finding group member by group id %v: %v", groupID, err)

		return nil, TranslateGormError(err)
	}

	return groupMembers, nil
}

func CreateGroupMember(tx *gorm.DB, member models.GroupMember) (err error) {
	if err = tx.Model(&models.GroupMember{}).Create(&member).Error; err != nil {
		logger.Error.Printf("[repository.CreateGroupMember] Error creating group member %v: %v", member, err)

		return TranslateGormError(err)
	}

	return nil
}

func UpdateGroupMember(member models.GroupMember) (err error) {
	if err = db.GetDBConn().Model(&models.GroupMember{}).Where("id = ?", member.ID).Updates(member).Error; err != nil {
		logger.Error.Printf("[repository.UpdateGroupMember] Error updating group member %v: %v", member, err)

		return TranslateGormError(err)
	}

	return nil
}

func DeleteGroupMember(memberID uint) (err error) {
	if err = db.GetDBConn().Model(&models.GroupMember{}).Where("id = ?", memberID).Delete(&models.GroupMember{}).Error; err != nil {
		logger.Error.Printf("[repository.DeleteGroupMember] Error deleting member: %v", err)

		return TranslateGormError(err)
	}

	return nil
}
