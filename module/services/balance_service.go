package services

import (
	"Master_Data/module/domain/master"
	"Master_Data/module/repository"
	"gorm.io/gorm"
)

type BalancesService struct {
	BalancesRepository *repository.BalancesRepository
}

func NewBalanceService(db *gorm.DB) *BalancesService {
	balancesRepo := repository.NewBalancesRepository(db)
	return &BalancesService{BalancesRepository: balancesRepo}
}

func (s BalancesService) CreateBalances(userID uint) (interface{}, error) {
	balances := master.Balance{
		UserID:  userID,
		Balance: 0.0,
	}
	err := s.BalancesRepository.CreateBalances(&balances)
	if err != nil {
		return nil, err
	}
	return &balances, nil
}

func (s BalancesService) GetBalancesByID(id uint) (interface{}, error) {
	return s.BalancesRepository.GetBalancesByID(id)
}

func (s BalancesService) GetBalancesByUserID(userID uint) (interface{}, error) {
	return s.BalancesRepository.GetBalancesByUserID(userID)
}

func (s BalancesService) UpdateBalancesBalance(id uint, balance *float64) (interface{}, error) {
	return s.BalancesRepository.UpdateBalancesBalance(id, balance)
}

func (s BalancesService) UpdateBalances(userID uint, balance *float64) (interface{}, error) {
	balances, err := s.GetBalancesByUserID(userID)
	if err != nil {
		return nil, err
	}
	balances, err = s.UpdateBalancesBalance(userID, balance)
	if err != nil {
		return nil, err
	}
	return balances, nil
}
