package repository

import (
	"technoCredits/internal/app/models"
	"technoCredits/pkg/db"
	"technoCredits/pkg/logger"
)

func GetAllUserGroups(userID uint) (group models.Group, err error) {
	if err = db.GetDBConn().
		Preload("GroupMembers").
		Preload("GroupMembers.Users").
		Joins("JOIN group_members ON group_members.group_id = cards_expenses.group_id").
		Where("group_members.user_id = ?", userID).Error; err != nil {
		logger.Error.Printf("[repository.GetAllUserGroups] Error when get group members: %v", err)

		return group, TranslateGormError(err)
	}

	return group, nil
}

//func CreateGroup(userID uint, group models.Group) (models.Group, error) {
//
//}
