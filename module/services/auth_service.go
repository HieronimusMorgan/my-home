package services

import (
	"Master_Data/module/domain"
	"Master_Data/module/dto/in"
	"Master_Data/module/repository"
	"Master_Data/utils"
	"gorm.io/gorm"
	"time"
)

type AuthService struct {
	UserRepository      *repository.UserRepository
	FinancialRepository *repository.FinancialRepository
	TokenRepository     *repository.TokenRepository
}

func NewAuthService(db *gorm.DB) *AuthService {
	userRepo := repository.NewUserRepository(db)
	financialRepo := repository.NewFinancialRepository(db)
	tokenRepo := repository.NewTokenRepository(db)
	return &AuthService{UserRepository: userRepo, FinancialRepository: financialRepo, TokenRepository: tokenRepo}
}

func (s AuthService) Register(i *in.RegisterRequest) (interface{}, error) {

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

	financial := domain.Financial{
		UserID:  newUser.(domain.User).UserID,
		Balance: 0.0,
	}
	err = s.FinancialRepository.CreateFinancial(&financial)
	if err != nil {
		return nil, err
	}

	token, err := utils.GenerateToken(user)
	if err != nil {
		return nil, err
	}
	mapToken := map[string]interface{}{
		"user":      newUser.(domain.User),
		"financial": financial,
		"token":     token,
	}
	utils.SaveTokenToRedis(newUser.(domain.User).ClientID, token, time.Hour*24)
	utils.SaveDataToRedis("user", newUser.(domain.User).ClientID, newUser)
	return mapToken, nil
}

func (s AuthService) Login(i *struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}) (interface{}, error) {
	user, err := s.UserRepository.GetUserByUsername(i.Username)
	if err != nil {
		return nil, err
	}

	err = utils.CheckPassword(user.(domain.User).Password, i.Password)
	if err != nil {
		return nil, err
	}

	financial, err := s.FinancialRepository.GetFinancialByUserID(user.(domain.User).UserID)
	if err != nil {
		return nil, err

	}
	token, err := utils.GenerateToken(user.(domain.User))
	if err != nil {
		return nil, err
	}
	refreshToken, err := utils.GenerateRefreshToken()
	s.TokenRepository.CreateToken(user.(domain.User), token, refreshToken)

	utils.SaveTokenToRedis(user.(domain.User).ClientID, token, time.Hour*24)
	utils.SaveDataToRedis("user", user.(domain.User).ClientID, user)

	mapToken := map[string]interface{}{
		"user":          utils.ConvertUserToResponse(user.(domain.User)),
		"balance":       financial.(domain.Financial).Balance,
		"token":         token,
		"refresh_token": refreshToken,
	}

	return mapToken, nil
}

func (s AuthService) RefreshToken(token, refreshToken string) (interface{}, error) {
	clientID, err := utils.GetClientIDFromToken(token)
	if err != nil {
		return nil, err
	}

	user, err := s.UserRepository.GetUserByClientID(clientID.(string))
	if err != nil {
		return nil, err
	}

	token, err = utils.GenerateToken(user.(domain.User))
	if err != nil {
		return nil, err
	}
	refreshToken, err = utils.GenerateRefreshToken()
	err = s.TokenRepository.RefreshToken(user.(domain.User), token, refreshToken)
	if err != nil {
		return nil, err
	}

	utils.SaveTokenToRedis(user.(domain.User).ClientID, token, time.Hour*24)
	utils.SaveDataToRedis("user", user.(domain.User).ClientID, user)

	mapToken := map[string]interface{}{
		"user":          utils.ConvertUserToResponse(user.(domain.User)),
		"token":         token,
		"refresh_token": refreshToken,
	}

	return mapToken, nil
}

func (s AuthService) Logout(token string) error {

	clientID, err := utils.GetClientIDFromToken(token)
	if err != nil {
		return err
	}

	user, err := s.UserRepository.GetUserByClientID(clientID.(string))
	if err != nil {
		return err
	}
	utils.DeleteTokenFromRedis(user.(domain.User).ClientID)

	err = s.TokenRepository.DeleteToken(user.(domain.User).UserID)
	if err != nil {
		return err
	}
	return nil
}
