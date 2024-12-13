package routes

import (
	"Master_Data/module/handler"
	"Master_Data/module/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BalancesRoutes(r *gin.Engine, db *gorm.DB) {
	// Initialize Handlers
	userHandler := handler.NewBalanceHandler(db)

	protected := r.Group("/api/v1/balances")
	protected.Use(middleware.MasterDataMiddleware()) // Apply middleware
	{
		protected.POST("/update", userHandler.UpdateBalance)
	}
}
