package repository

import (
	"gorm.io/gorm"
	"technoCredits/internal/app/models"
	"technoCredits/pkg/logger"
)

type SettlementRepository struct {
	db *gorm.DB
}

func NewSettlementRepository(db *gorm.DB) *SettlementRepository {
	return &SettlementRepository{db: db}
}

func (r *SettlementRepository) Create(settlement *models.Settlement) error {
	if err := r.db.Create(settlement).Error; err != nil {
		logger.Error.Printf("[repository.Settlement.Create] error: %v", err)
		return TranslateGormError(err)
	}
	logger.Info.Printf("[repository.Settlement.Create] created settlement ID=%d", settlement.ID)
	return nil
}

func (r *SettlementRepository) GetByID(id uint) (*models.Settlement, error) {
	var settlement models.Settlement
	if err := r.db.First(&settlement, id).Error; err != nil {
		logger.Error.Printf("[repository.Settlement.GetByID] error: %v", err)
		return nil, TranslateGormError(err)
	}
	return &settlement, nil
}

func (r *SettlementRepository) GetAll() ([]models.Settlement, error) {
	var settlements []models.Settlement
	if err := r.db.Find(&settlements).Error; err != nil {
		logger.Error.Printf("[repository.Settlement.GetAll] error: %v", err)
		return nil, TranslateGormError(err)
	}
	return settlements, nil
}

func (r *SettlementRepository) Update(settlement *models.Settlement) error {
	if err := r.db.Save(settlement).Error; err != nil {
		logger.Error.Printf("[repository.Settlement.Update] error: %v", err)
		return TranslateGormError(err)
	}
	logger.Info.Printf("[repository.Settlement.Update] updated settlement ID=%d", settlement.ID)
	return nil
}

func (r *SettlementRepository) Delete(id uint) error {
	if err := r.db.Delete(&models.Settlement{}, id).Error; err != nil {
		logger.Error.Printf("[repository.Settlement.Delete] error: %v", err)
		return TranslateGormError(err)
	}
	logger.Info.Printf("[repository.Settlement.Delete] deleted settlement ID=%d", id)
	return nil
}
