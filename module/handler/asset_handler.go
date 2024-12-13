package handler

import (
	"Master_Data/module/services"
	"Master_Data/package/response"
	"Master_Data/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AssetHandler struct {
	AssetService *services.AssetService
}

func NewAssetHandler(db *gorm.DB) *AssetHandler {
	s := services.NewAssetService(db)
	return &AssetHandler{AssetService: s}
}

func (h AssetHandler) AddAssetCategory(context *gin.Context) {
	var req struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description" binding:"required"`
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
	assetCategory, err := h.AssetService.AddAssetCategory(&req.Name, &req.Description, clientID.(string))
	if err != nil {
		response.SendResponse(context, 500, "Failed to add asset category", nil, err)
		return
	}

	response.SendResponse(context, 200, "Asset category added successfully", assetCategory, nil)
}

func (h AssetHandler) GetAssetCategoryByName(context *gin.Context) {
	name := context.Param("id")
	assetCategory, err := h.AssetService.GetAssetCategoryByName(&name)
	if err != nil {
		response.SendResponse(context, 500, "Failed to get asset category", nil, err)
		return
	}

	response.SendResponse(context, 200, "Asset category retrieved successfully", assetCategory, nil)
}

func (h AssetHandler) GetAssetCategoryById(context *gin.Context) {
	id := context.Param("id")
	assetCategory, err := h.AssetService.GetAssetCategoryById(&id)
	if err != nil {
		response.SendResponse(context, 500, "Failed to get asset category", nil, err)
		return
	}

	response.SendResponse(context, 200, "Asset category retrieved successfully", assetCategory, nil)
}
