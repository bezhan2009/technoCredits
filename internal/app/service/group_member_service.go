package service

import (
	"technoCredits/internal/app/models"
	"technoCredits/internal/repository"
	"technoCredits/pkg/db"
)

func GetGroupMembersByGroupID(groupID uint) (groupMembers models.GroupMember, err error) {
	return repository.GetGroupMembersByGroupID(groupID)
}

func CreateGroupMember(member models.GroupMember) (err error) {
	tx := db.GetDBConn().Begin()
	err = repository.CreateGroupMember(tx, member)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func UpdateGroupMember(member models.GroupMember) (err error) {
	return repository.UpdateGroupMember(member)
}

func DeleteGroupMember(memberID uint) (err error) {
	return repository.DeleteGroupMember(memberID)
}
