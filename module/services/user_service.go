package services

import (
	"Master_Data/module/domain"
	"Master_Data/module/dto/in"
	"Master_Data/module/repository"
	"Master_Data/utils"
	"gorm.io/gorm"
	"time"
)

type UserService struct {
	UserRepository      *repository.UserRepository
	FinancialRepository *repository.FinancialRepository
}

func NewUserService(db *gorm.DB) *UserService {
	userRepo := repository.NewUserRepository(db)
	financialRepo := repository.NewFinancialRepository(db)
	return &UserService{UserRepository: userRepo, FinancialRepository: financialRepo}
}

func (s UserService) CreateUser(i *in.RegisterRequest) (interface{}, error) {
	if err := utils.ValidateUsername(i.Username); err != nil {
		return nil, err
	}
	pass, err := utils.HashPassword(i.Password)
	if err != nil {
		return nil, err
	}

	firstName := utils.ValidationTrimSpace(i.FirstName)
	lastName := utils.ValidationTrimSpace(i.LastName)
	fullName := firstName + " " + lastName

	user := domain.User{
		Username:       i.Username,
		Password:       pass,
		FirstName:      firstName,
		LastName:       lastName,
		FullName:       fullName,
		PhoneNumber:    i.PhoneNumber,
		ProfilePicture: i.ProfilePicture,
		RoleID:         2,
	}
	err = s.UserRepository.CreateUser(&user)
	if err != nil {
		return nil, err
	}
	newUser, err := s.UserRepository.GetUserByUsername(i.Username)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	token, err := utils.GenerateToken(user)
	if err != nil {
		return nil, err
	}
	mapToken := map[string]interface{}{
		"user":  newUser.(domain.User),
		"token": token,
	}
	utils.SaveTokenToRedis(newUser.(domain.User).ClientID, token, time.Hour*24)
	utils.SaveDataToRedis("user", newUser.(domain.User).ClientID, newUser)
	return mapToken, nil
}

func (s UserService) GetUserProfile(clientID string) (interface{}, error) {
	user, err := s.UserRepository.GetUserByClientID(clientID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
