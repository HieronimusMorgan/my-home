package repository

import (
	"Master_Data/module/domain"
	"gorm.io/gorm"
)

type FinancialRepository struct {
	DB *gorm.DB
}

func NewFinancialRepository(db *gorm.DB) *FinancialRepository {
	return &FinancialRepository{DB: db}
}

func (r *FinancialRepository) CreateFinancial(financial *domain.Financial) error {
	return r.DB.Create(financial).Error
}

func (r *FinancialRepository) GetFinancialByID(id uint) (interface{}, error) {
	var financial domain.Financial
	err := r.DB.First(&financial, id).Error
	return financial, err
}

func (r *FinancialRepository) GetFinancialByUserID(userID uint) (interface{}, error) {
	var financial domain.Financial
	err := r.DB.Where("user_id = ?", userID).First(&financial).Error
	return financial, err
}

func (r *FinancialRepository) UpdateFinancialBalance(userID uint, balance float64) error {
	tx := r.DB.Begin()
	if err := tx.Model(&domain.Financial{}).Where("user_id = ?", userID).Update("balance", balance).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (r *FinancialRepository) DeleteFinancial(userID uint) error {
	tx := r.DB.Begin()
	if err := tx.Where("user_id = ?", userID).Delete(&domain.Financial{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
