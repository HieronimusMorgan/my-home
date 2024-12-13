package repository

import (
	"Master_Data/module/domain/master"
	"fmt"
	"gorm.io/gorm"
)

type TokenRepository struct {
	DB *gorm.DB
}

func NewTokenRepository(db *gorm.DB) *TokenRepository {
	return &TokenRepository{DB: db}
}

func (r TokenRepository) CreateToken(user master.User, token, refreshToken string) error {
	tx := r.DB.Begin()
	if err := tx.Error; err != nil {
		return err
	}
	var existingToken master.Token
	if err := tx.Where("user_id = ?", user.UserID).First(&existingToken).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			newToken := master.Token{
				UserID:       user.UserID,
				Token:        token,
				RefreshToken: refreshToken,
				CreatedBy:    user.ClientID,
				UpdatedBy:    user.ClientID,
			}
			if err := tx.Create(&newToken).Error; err != nil {
				tx.Rollback()
				return err
			}
		} else {
			tx.Rollback()
			return err
		}
	} else {
		if err := tx.Model(&existingToken).
			Updates(map[string]interface{}{
				"token":         token,
				"refresh_token": refreshToken,
				"updated_by":    user.ClientID,
			}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

func (r TokenRepository) GetToken(userID uint) (interface{}, error) {
	var token master.Token
	err := r.DB.Where("user_id = ?", userID).First(&token).Error
	return token, err
}

func (r TokenRepository) RefreshToken(user master.User, token, refreshToken string) error {
	tx := r.DB.Begin()
	if err := tx.Error; err != nil {
		return err
	}
	var existingToken master.Token
	if err := tx.Where("user_id = ?", user.UserID).First(&existingToken).Error; err == nil {
		if err := tx.Model(&existingToken).
			Updates(map[string]interface{}{
				"token":         token,
				"refresh_token": refreshToken,
				"updated_by":    user.ClientID,
			}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

func (r TokenRepository) DeleteToken(userID uint) error {
	tx := r.DB.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	var token master.Token
	if err := tx.Where("user_id = ?", userID).First(&token).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			tx.Rollback()
			return fmt.Errorf("token not found for user_id: %s", userID)
		}
		tx.Rollback()
		return err
	}

	if err := tx.Model(&token).Update("expired", true).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
