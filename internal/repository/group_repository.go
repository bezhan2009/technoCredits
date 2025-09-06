package repository

import (
	"gorm.io/gorm"
	"technoCredits/internal/app/models"
	"technoCredits/pkg/db"
	"technoCredits/pkg/logger"
)

func GetAllUserGroups(userID uint) ([]models.Group, error) {
	var groups []models.Group
	err := db.GetDBConn().
		Preload("GroupMembers").
		Preload("GroupMembers.User").
		Joins("JOIN group_members ON group_members.group_id = groups.id").
		Where("group_members.user_id = ?", userID).
		Find(&groups).Error

	if err != nil {
		logger.Error.Printf("[repository.GetAllUserGroups] Error when get group members: %v", err)
		return nil, TranslateGormError(err)
	}

	return groups, nil
}

func GetAllUserGroupsByID(userID uint, id uint) (models.Group, error) {
	var group models.Group
	err := db.GetDBConn().
		Preload("GroupMembers").
		Preload("GroupMembers.User").
		Joins("JOIN group_members ON group_members.group_id = groups.id").
		Where("group_members.user_id = ? AND groups.id = ?", userID, id).
		First(&group).Error

	if err != nil {
		logger.Error.Printf("[repository.GetAllUserGroupsByID] Error when get group members: %v", err)
		return models.Group{}, TranslateGormError(err)
	}

	return group, nil
}

func CreateGroup(tx *gorm.DB, group models.Group) (id uint, err error) {
	if err = tx.Model(&models.Group{}).Create(&group).Error; err != nil {
		logger.Error.Printf("[repository.CreateGroup] Error when create group: %v", err)

		return 0, TranslateGormError(err)
	}

	return group.ID, nil
}

func UpdateGroup(group models.Group) (err error) {
	if err = db.GetDBConn().Model(&models.Group{}).Where("id = ?", group.ID).Updates(&group).Error; err != nil {
		logger.Error.Printf("[repository.UpdateGroup] Error when update group: %v", err)

		return TranslateGormError(err)
	}

	return nil
}

func DeleteGroup(groupID uint) (err error) {
	if err = db.GetDBConn().Model(&models.Group{}).Where("id = ?", groupID).Delete(&models.Group{}).Error; err != nil {
		logger.Error.Printf("[repository.DeleteGroup] Error when delete group: %v", err)

		return TranslateGormError(err)
	}

	return nil
}
