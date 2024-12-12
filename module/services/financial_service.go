package services

import (
	"Master_Data/module/domain"
	"Master_Data/module/repository"
)

type FinancialService struct {
	FinancialRepository *repository.FinancialRepository
}

func NewFinancialService(financialRepo *repository.FinancialRepository) *FinancialService {
	return &FinancialService{FinancialRepository: financialRepo}
}

func (s FinancialService) CreateFinancial(userID uint) (interface{}, error) {
	financial := domain.Financial{
		UserID:  userID,
		Balance: 0.0,
	}
	err := s.FinancialRepository.CreateFinancial(&financial)
	if err != nil {
		return nil, err
	}
	return &financial, nil
}

func (s FinancialService) GetFinancialByID(id uint) (interface{}, error) {
	return s.FinancialRepository.GetFinancialByID(id)
}

func (s FinancialService) GetFinancialByUserID(userID uint) (interface{}, error) {
	return s.FinancialRepository.GetFinancialByUserID(userID)
}

func (s FinancialService) UpdateFinancialBalance(id uint, balance float64) error {
	return s.FinancialRepository.UpdateFinancialBalance(id, balance)
}
