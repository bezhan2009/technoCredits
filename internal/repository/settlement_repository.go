package repository

import (
	"technoCredits/internal/app/models"
	"technoCredits/pkg/db"
	"technoCredits/pkg/logger"
)

func SettlementCreate(settlement *models.Settlement) error {
	if err := db.GetDBConn().Create(settlement).Error; err != nil {
		logger.Error.Printf("[repository.SettlementCreate] error: %v", err)

		return TranslateGormError(err)
	}

	logger.Info.Printf("[repository.SettlementCreate] created settlement ID=%d", settlement.ID)

	return nil
}

func GetSettlementByID(id uint) (*models.Settlement, error) {
	var settlement models.Settlement
	if err := db.GetDBConn().Preload("FromUser").
		First(&settlement, id).Error; err != nil {
		logger.Error.Printf("[repository.GetSettlementByID] error: %v", err)

		return nil, TranslateGormError(err)
	}

	return &settlement, nil
}

func GetAllSettlements() ([]models.Settlement, error) {
	var settlements []models.Settlement
	if err := db.GetDBConn().Preload("FromUser").Find(&settlements).Error; err != nil {
		logger.Error.Printf("[repository.GetAllSettlements] error: %v", err)

		return nil, TranslateGormError(err)
	}

	return settlements, nil
}

func UpdateSettlements(settlement *models.Settlement) error {
	if err := db.GetDBConn().Where("id = ?", settlement.ID).Updates(settlement).Error; err != nil {
		logger.Error.Printf("[repository.UpdateSettlements] error: %v", err)

		return TranslateGormError(err)
	}

	logger.Info.Printf("[repository.UpdateSettlements] updated settlement ID=%d", settlement.ID)

	return nil
}

func DeleteSettlement(id uint) error {
	if err := db.GetDBConn().Delete(&models.Settlement{}, id).Error; err != nil {
		logger.Error.Printf("[repository.DeleteSettlement] error: %v", err)

		return TranslateGormError(err)
	}

	logger.Info.Printf("[repository.DeleteSettlement] deleted settlement ID=%d", id)

	return nil
}
