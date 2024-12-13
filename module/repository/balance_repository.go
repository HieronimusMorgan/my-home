package repository

import (
	"Master_Data/module/domain/master"
	"gorm.io/gorm"
)

type BalancesRepository struct {
	DB *gorm.DB
}

func NewBalancesRepository(db *gorm.DB) *BalancesRepository {
	return &BalancesRepository{DB: db}
}

func (r *BalancesRepository) CreateBalances(balances *master.Balance) error {
	return r.DB.Create(balances).Error
}

func (r *BalancesRepository) GetBalancesByID(id uint) (interface{}, error) {
	var balances master.Balance
	err := r.DB.First(&balances, id).Error
	return balances, err
}

func (r *BalancesRepository) GetBalancesByUserID(userID uint) (interface{}, error) {
	var balances master.Balance
	err := r.DB.Where("user_id = ?", userID).First(&balances).Error
	return balances, err
}

func (r *BalancesRepository) UpdateBalancesBalance(userID uint, balance *float64) (interface{}, error) {
	var balances master.Balance
	tx := r.DB.Begin()
	if err := tx.Model(&balances).Where("user_id = ?", userID).Update("balance", balance).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return balances, nil
}

func (r *BalancesRepository) DeleteBalances(userID uint) error {
	tx := r.DB.Begin()
	if err := tx.Where("user_id = ?", userID).Delete(&master.Balance{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
