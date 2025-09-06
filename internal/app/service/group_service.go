package service

import (
	"errors"
	_ "log"
	"technoCredits/internal/app/models"
	"technoCredits/internal/repository"
)

type GroupService struct {
	repo *repository.GroupRepository
}

func NewGroupService(repo *repository.GroupRepository) *GroupService {
	return &GroupService{repo: repo}
}

func (s *GroupService) Create(group *models.Group) error {
	if group.Name == "" {
		return errors.New("название группы не может быть пустым")
	}
	return s.repo.Create(group)
}

func (s *GroupService) GetByID(id uint) (*models.Group, error) {
	return s.repo.GetByID(id)
}

func (s *GroupService) Update(group *models.Group) error {
	if group.ID == 0 {
		return errors.New("ID группы не может быть нулевым")
	}
	return s.repo.Update(group)
}

func (s *GroupService) Delete(id uint) error {
	return s.repo.Delete(id)
}
