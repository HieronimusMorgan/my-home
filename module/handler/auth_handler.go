package handler

import (
	"Master_Data/module/dto/in"
	"Master_Data/module/services"
	"Master_Data/package/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type AuthHandler struct {
	AuthService *services.AuthService
}

func NewAuthHandler(db *gorm.DB) *AuthHandler {
	s := services.NewAuthService(db)
	return &AuthHandler{AuthService: s}
}

func (h AuthHandler) Register(c *gin.Context) {
	var req in.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendResponse(c, http.StatusBadRequest, "Error", nil, err.Error())
		return
	}

	user, err := h.AuthService.Register(&req)
	if err != nil {
		response.SendResponse(c, http.StatusInternalServerError, "Failed to register user", nil, err)
		return
	}

	response.SendResponse(c, http.StatusCreated, "User registered successfully", user, nil)
}

func (h AuthHandler) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendResponse(c, http.StatusBadRequest, "Error", nil, err.Error())
		return
	}

	user, err := h.AuthService.Login(&req)
	if err != nil {
		response.SendResponse(c, http.StatusInternalServerError, "Failed to login", nil, err)
		return
	}

	response.SendResponse(c, http.StatusOK, "Login success", user, nil)
}

func (h AuthHandler) RefreshToken(c *gin.Context) {
	token := c.GetHeader("Authorization")

	var req struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendResponse(c, http.StatusBadRequest, "Error", nil, err.Error())
		return
	}

	if token == "" {
		response.SendResponse(c, http.StatusBadRequest, "Error", nil, "Refresh token is required")
		return
	}

	user, err := h.AuthService.RefreshToken(token, req.RefreshToken)
	if err != nil {
		response.SendResponse(c, http.StatusInternalServerError, "Failed to refresh token", nil, err)
		return
	}

	response.SendResponse(c, http.StatusOK, "Token refreshed", user, nil)

}

func (h AuthHandler) Logout(c *gin.Context) {
	token := c.GetHeader("Authorization")
	err := h.AuthService.Logout(token)
	if err != nil {
		response.SendResponse(c, http.StatusInternalServerError, "Failed to logout", nil, err)
		return
	}

	response.SendResponse(c, http.StatusOK, "Logout success", nil, nil)
}
