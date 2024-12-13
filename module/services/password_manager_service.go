package services

import (
	"Master_Data/module/domain/master"
	"Master_Data/module/repository"
	"Master_Data/utils"
	"gorm.io/gorm"
)

type PasswordManagerService struct {
	PasswordManagerRepository *repository.PasswordManagerRepository
	UserRepository            *repository.UserRepository
}

func NewPasswordManagerService(db *gorm.DB) *PasswordManagerService {
	balancesRepo := repository.NewPasswordManagerRepository(db)
	userRepo := repository.NewUserRepository(db)
	return &PasswordManagerService{PasswordManagerRepository: balancesRepo, UserRepository: userRepo}
}

func (s PasswordManagerService) AddPassword(name string, password *string, description string, clientID string) (interface{}, error) {
	user, err := s.UserRepository.GetUserByClientID(clientID)
	if err != nil {
		return nil, err
	}
	err = utils.ValidatePassword(password)

	passwordEncrypt, err := utils.EncryptAES(password, user.(master.User).ClientID)

	passwordManager := master.PasswordManager{
		UserID:      user.(master.User).UserID,
		Name:        name,
		Password:    passwordEncrypt,
		Description: description,
		Expired:     false,
		CreatedBy:   user.(master.User).ClientID,
		UpdatedBy:   user.(master.User).ClientID,
	}

	err = s.PasswordManagerRepository.AddPassword(&passwordManager)
	if err != nil {
		return nil, err
	}

	return passwordManager.Name, nil
}
