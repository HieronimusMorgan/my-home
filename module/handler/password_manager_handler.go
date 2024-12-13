package handler

import (
	"Master_Data/module/services"
	"Master_Data/package/response"
	"Master_Data/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PasswordManagerHandler struct {
	PasswordManagerService *services.PasswordManagerService
}

func NewPasswordManagerHandler(db *gorm.DB) *PasswordManagerHandler {
	s := services.NewPasswordManagerService(db)
	return &PasswordManagerHandler{PasswordManagerService: s}
}

func (h PasswordManagerHandler) AddPassword(context *gin.Context) {
	var req struct {
		Name        string `json:"name" binding:"required"`
		Password    string `json:"password" binding:"required"`
		Description string `json:"description" binding:"required"`
		Url         string `json:"url" binding:"omitempty"`
	}
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}
	clientID, err := utils.GetClientIDFromToken(context.GetHeader("Authorization"))
	if err != nil {
		response.SendResponse(context, 500, "Failed to get user id from token", nil, err)
		return
	}
	passwordManager, err := h.PasswordManagerService.AddPassword(req.Name, &req.Password, req.Description, clientID.(string))
	if err != nil {
		response.SendResponse(context, 400, "Failed to add password", nil, err)
		return
	}

	response.SendResponse(context, 200, "Password added successfully", passwordManager, nil)
}

func (h PasswordManagerHandler) UpdatePassword(context *gin.Context) {

}

func (h PasswordManagerHandler) DeletePassword(context *gin.Context) {

}

func (h PasswordManagerHandler) GetPassword(context *gin.Context) {
	//name := context.Param("name")
	//clientID, err := utils.GetClientIDFromToken(context.GetHeader("Authorization"))
	//if err != nil {
	//	response.SendResponse(context, 500, "Failed to get user id from token", nil, err)
	//	return
	//}
	//passwordManager, err := h.PasswordManagerService.GetPassword(name, clientID.(string))

}
