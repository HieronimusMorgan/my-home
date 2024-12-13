package handler

import (
	"Master_Data/module/domain/master"
	"Master_Data/module/services"
	"Master_Data/package/response"
	"Master_Data/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type BalancesHandler struct {
	BalancesService *services.BalancesService
}

func NewBalanceHandler(db *gorm.DB) *BalancesHandler {
	s := services.NewBalanceService(db)
	return &BalancesHandler{BalancesService: s}
}

func (h BalancesHandler) UpdateBalance(c *gin.Context) {
	var req struct {
		Balance float64 `json:"balance" binding:"required"`
	}

	token := c.GetHeader("Authorization")

	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendResponse(c, http.StatusBadRequest, "Error", nil, err.Error())
		return
	}
	userID, err := utils.GetUserIDFromToken(token)
	if err != nil {
		response.SendResponse(c, http.StatusUnauthorized, "Failed to get user ID from token", nil, err)
	}
	balances, err := h.BalancesService.UpdateBalances(userID, &req.Balance)
	if err != nil {
		response.SendResponse(c, http.StatusInternalServerError, "Failed to update balances", nil, err)
		return
	}
	mapBalances := map[string]interface{}{
		"balance": balances.(master.Balance).Balance,
	}
	response.SendResponse(c, http.StatusOK, "Balance updated successfully", mapBalances, nil)
}
