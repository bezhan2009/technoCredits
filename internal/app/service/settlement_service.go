package service

import (
	"technoCredits/internal/app/models"
	"technoCredits/internal/repository"
	"technoCredits/pkg/logger"
)

type SettlementService struct {
	repo *repository.SettlementRepository
}

func NewSettlementService(repo *repository.SettlementRepository) *SettlementService {
	return &SettlementService{repo: repo}
}

func (s *SettlementService) Create(settlement *models.Settlement) error {
	logger.Info.Printf("[service.Settlement.Create] creating settlement for GroupID=%d", settlement.GroupID)
	return s.repo.Create(settlement)
}

func (s *SettlementService) GetByID(id uint) (*models.Settlement, error) {
	logger.Info.Printf("[service.Settlement.GetByID] fetching settlement ID=%d", id)
	return s.repo.GetByID(id)
}

func (s *SettlementService) GetAll() ([]models.Settlement, error) {
	logger.Info.Println("[service.Settlement.GetAll] fetching all settlements")
	return s.repo.GetAll()
}

func (s *SettlementService) Update(settlement *models.Settlement) error {
	logger.Info.Printf("[service.Settlement.Update] updating settlement ID=%d", settlement.ID)
	return s.repo.Update(settlement)
}

func (s *SettlementService) Delete(id uint) error {
	logger.Info.Printf("[service.Settlement.Delete] deleting settlement ID=%d", id)
	return s.repo.Delete(id)
}
