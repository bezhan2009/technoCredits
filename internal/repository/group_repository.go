package repository

import (
	"gorm.io/gorm"
	"log"
	"technoCredits/internal/app/models"
)

type GroupRepository struct {
	DB *gorm.DB
}

func NewGroupRepository(db *gorm.DB) *GroupRepository {
	return &GroupRepository{DB: db}
}

func (r *GroupRepository) Create(group *models.Group) error {
	log.Printf("Создание группы: %s", group.Name)
	return TranslateGormError(r.DB.Create(group).Error)
}

func (r *GroupRepository) GetByID(id uint) (*models.Group, error) {
	var group models.Group
	if err := r.DB.Preload("Owner").First(&group, id).Error; err != nil {
		return nil, TranslateGormError(err)
	}
	return &group, nil
}

func (r *GroupRepository) Update(group *models.Group) error {
	return TranslateGormError(r.DB.Save(group).Error)
}

func (r *GroupRepository) Delete(id uint) error {
	return TranslateGormError(r.DB.Delete(&models.Group{}, id).Error)
}
