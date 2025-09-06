package service

import (
	"errors"
	"technoCredits/internal/app/models"
	"technoCredits/internal/repository"
	"time"
)

var (
	ErrAlreadyInGroup = errors.New("user is already a member of the group")
	ErrNotFound       = errors.New("group member not found")
)

type GroupMemberService struct {
	repo *repository.GroupMemberRepository
}

func NewGroupMemberService(repo *repository.GroupMemberRepository) *GroupMemberService {
	return &GroupMemberService{repo: repo}
}

func (s *GroupMemberService) AddMember(groupID uint, userID uint, role string) error {
	members, err := s.repo.GetByGroupID(groupID)
	if err != nil {
		return err
	}
	for _, m := range members {
		if m.UserID == userID {
			return ErrAlreadyInGroup
		}
	}

	member := &models.GroupMember{
		GroupID:  groupID,
		UserID:   userID,
		Role:     role,
		JoinedAt: time.Now(),
	}

	return s.repo.Create(member)
}

func (s *GroupMemberService) GetMembers(groupID uint) ([]models.GroupMember, error) {
	return s.repo.GetByGroupID(groupID)
}

func (s *GroupMemberService) UpdateMemberRole(groupID uint, userID uint, role string) error {
	return s.repo.UpdateRole(groupID, userID, role)
}

func (s *GroupMemberService) RemoveMember(groupID uint, userID uint) error {
	members, err := s.repo.GetByGroupID(groupID)
	if err != nil {
		return err
	}

	found := false
	for _, m := range members {
		if m.UserID == userID {
			found = true
			break
		}
	}
	if !found {
		return ErrNotFound
	}

	return s.repo.Delete(groupID, userID)
}
