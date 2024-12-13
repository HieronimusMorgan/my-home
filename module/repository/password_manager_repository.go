package repository

import (
	"Master_Data/module/domain/master"
	"gorm.io/gorm"
)

type PasswordManagerRepository struct {
	DB *gorm.DB
}

func NewPasswordManagerRepository(db *gorm.DB) *PasswordManagerRepository {
	return &PasswordManagerRepository{DB: db}
}

func (r PasswordManagerRepository) AddPassword(passwordManager *master.PasswordManager) error {
	return r.DB.Create(passwordManager).Error
}
