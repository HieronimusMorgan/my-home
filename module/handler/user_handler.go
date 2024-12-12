package handler

import (
	"Master_Data/module/dto/in"
	"Master_Data/module/services"
	"Master_Data/package/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	s := services.NewUserService(db) // Initialize Service Layer
	return &UserHandler{UserService: s}
}

func (h UserHandler) RegisterUser(c *gin.Context) {
	var req in.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendResponse(c, http.StatusBadRequest, "Error", nil, err.Error())
		return
	}

	user, err := h.UserService.CreateUser(&req)
	if err != nil {
		response.SendResponse(c, http.StatusInternalServerError, "Failed to register user", nil, err)
		return
	}

	response.SendResponse(c, http.StatusCreated, "User registered successfully", user, nil)
}

func (h UserHandler) GetUserProfile(c *gin.Context) {
	clientID, exists := c.Get("client_id")
	if !exists {
		response.SendResponse(c, http.StatusUnauthorized, "User ID not found in context", nil, "Missing user ID")
		return
	}
	clientIDString, ok := clientID.(string)
	if !ok {
		response.SendResponse(c, http.StatusInternalServerError, "Failed to process user ID", nil, "Invalid user ID type")
		return
	}
	user, err := h.UserService.GetUserProfile(clientIDString)
	if err != nil {
		response.SendResponse(c, http.StatusInternalServerError, "Failed to get user profile", nil, err.Error())
		return
	}
	response.SendResponse(c, http.StatusOK, "User profile retrieved successfully", user, nil)
}
