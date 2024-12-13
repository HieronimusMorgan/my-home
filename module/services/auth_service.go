package services

import (
	"Master_Data/module/domain/master"
	"Master_Data/module/dto/in"
	"Master_Data/module/dto/out"
	"Master_Data/module/repository"
	"Master_Data/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type AuthService struct {
	UserRepository     *repository.UserRepository
	BalancesRepository *repository.BalancesRepository
	TokenRepository    *repository.TokenRepository
}

func NewAuthService(db *gorm.DB) *AuthService {
	userRepo := repository.NewUserRepository(db)
	balancesRepo := repository.NewBalancesRepository(db)
	tokenRepo := repository.NewTokenRepository(db)
	return &AuthService{UserRepository: userRepo, BalancesRepository: balancesRepo, TokenRepository: tokenRepo}
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
	userUUID := uuid.New().String()

	user := master.User{
		UUIDKey:        userUUID,
		ClientID:       utils.GenerateClientID(),
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

	balances := master.Balance{
		UserID:  newUser.(master.User).UserID,
		Balance: 0.0,
	}
	err = s.BalancesRepository.CreateBalances(&balances)
	if err != nil {
		return nil, err
	}

	token, err := utils.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	response := out.RegisterResponse{
		UserID:         newUser.(master.User).UserID,
		Username:       newUser.(master.User).Username,
		FirstName:      newUser.(master.User).FirstName,
		LastName:       newUser.(master.User).LastName,
		PhoneNumber:    newUser.(master.User).PhoneNumber,
		ProfilePicture: newUser.(master.User).ProfilePicture,
		Balance:        balances.Balance,
		Token:          token,
	}
	utils.SaveTokenToRedis(newUser.(master.User).ClientID, token, time.Hour*24)
	utils.SaveDataToRedis("user", newUser.(master.User).ClientID, newUser)
	return response, nil
}

func (s AuthService) Login(i *struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}) (interface{}, error) {
	user, err := s.UserRepository.GetUserByUsername(i.Username)
	if err != nil {
		return nil, err
	}

	err = utils.CheckPassword(user.(master.User).Password, i.Password)
	if err != nil {
		return nil, err
	}

	balances, err := s.BalancesRepository.GetBalancesByUserID(user.(master.User).UserID)

	if err != nil {
		return nil, err

	}
	token, err := utils.GenerateToken(user.(master.User))
	if err != nil {
		return nil, err
	}
	refreshToken, err := utils.GenerateRefreshToken()
	s.TokenRepository.CreateToken(user.(master.User), token, refreshToken)

	utils.SaveTokenToRedis(user.(master.User).ClientID, token, time.Hour*24)
	utils.SaveDataToRedis("user", user.(master.User).ClientID, user)

	loginResponse := out.LoginResponse{
		UserID:         user.(master.User).UserID,
		ClientID:       user.(master.User).ClientID,
		Username:       user.(master.User).Username,
		FirstName:      user.(master.User).FirstName,
		LastName:       user.(master.User).LastName,
		PhoneNumber:    user.(master.User).PhoneNumber,
		ProfilePicture: user.(master.User).ProfilePicture,
		Balance:        balances.(master.Balance).Balance,
		Token:          token,
		RefreshToken:   refreshToken,
	}

	return loginResponse, nil
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

	token, err = utils.GenerateToken(user.(master.User))
	if err != nil {
		return nil, err
	}
	refreshToken, err = utils.GenerateRefreshToken()
	err = s.TokenRepository.RefreshToken(user.(master.User), token, refreshToken)
	if err != nil {
		return nil, err
	}

	utils.SaveTokenToRedis(user.(master.User).ClientID, token, time.Hour*24)
	utils.SaveDataToRedis("user", user.(master.User).ClientID, user)

	mapToken := map[string]interface{}{
		"user":          utils.ConvertUserToResponse(user.(master.User)),
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
	utils.DeleteTokenFromRedis(user.(master.User).ClientID)

	err = s.TokenRepository.DeleteToken(user.(master.User).UserID)
	if err != nil {
		return err
	}
	return nil
}
