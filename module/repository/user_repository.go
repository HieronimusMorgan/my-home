package repository

import (
	"Master_Data/module/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r UserRepository) CreateUser(user *domain.User) error {
	return r.DB.Create(user).Error
}

func (r UserRepository) GetUserByUsername(username string) (interface{}, error) {
	var user domain.User
	err := r.DB.Where("username = ?", username).First(&user).Error
	return user, err
}

func (r UserRepository) GetUserByClientID(id string) (interface{}, error) {
	var user domain.User
	err := r.DB.Where("client_id = ?", id).First(&user).Error
	return user, err
}

func (r UserRepository) GetUserByID(id uint) (interface{}, error) {
	var user domain.User
	err := r.DB.First(&user, id).Error
	return user, err
}

func (r UserRepository) UpdateUser(user *domain.User) error {
	tc := r.DB.Begin()
	if err := tc.Save(user).Error; err != nil {
		tc.Rollback()
		return err
	}
	tc.Commit()
	return nil
}

func (r UserRepository) DeleteUser(id uint) error {
	tc := r.DB.Begin()
	if err := tc.Delete(&domain.User{}, id).Error; err != nil {
		tc.Rollback()
		return err
	}
	tc.Commit()
	return nil
}
