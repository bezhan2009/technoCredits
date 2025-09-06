package service

import (
	"technoCredits/internal/app/models"
	"technoCredits/internal/repository"
)

// SettlementCreate создает новое расчетное соглашение
func SettlementCreate(settlement *models.Settlement) error {
	return repository.SettlementCreate(settlement)
}

// GetSettlementByID возвращает расчетное соглашение по ID
func GetSettlementByID(id uint) (*models.Settlement, error) {
	return repository.GetSettlementByID(id)
}

// GetAllSettlements возвращает все расчетные соглашения
func GetAllSettlements() ([]models.Settlement, error) {
	return repository.GetAllSettlements()
}

// UpdateSettlements обновляет расчетное соглашение
func UpdateSettlements(settlement *models.Settlement) error {
	return repository.UpdateSettlements(settlement)
}

// DeleteSettlement удаляет расчетное соглашение по ID
func DeleteSettlement(id uint) error {
	return repository.DeleteSettlement(id)
}
