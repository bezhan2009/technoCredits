package service

import (
	_ "log"
	"technoCredits/internal/app/models"
	"technoCredits/internal/repository"
	"technoCredits/pkg/db"
)

func GetAllUserGroups(userID uint) ([]models.Group, error) {
	return repository.GetAllUserGroups(userID)
}

func GetAllUserGroupsByID(userID uint, id uint) (group models.Group, err error) {
	return repository.GetAllUserGroupsByID(userID, id)
}

func CreateGroup(group models.Group) (err error) {
	tx := db.GetDBConn().Begin()

	id, err := repository.CreateGroup(tx, group)
	if err != nil {
		tx.Rollback()

		return err
	}

	err = repository.CreateGroupMember(tx, models.GroupMember{
		GroupID: id,
		UserID:  group.OwnerID,
		RoleID:  1,
	})
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func UpdateGroup(group models.Group) (err error) {
	return repository.UpdateGroup(group)
}

func DeleteGroup(groupID uint) (err error) {
	return repository.DeleteGroup(groupID)
}
