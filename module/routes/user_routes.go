package routes

import (
	"Master_Data/module/handler"
	"Master_Data/module/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.Engine, db *gorm.DB) {
	// Initialize Handlers
	userHandler := handler.NewUserHandler(db)

	protected := r.Group("/api/v1/user")
	protected.Use(middleware.MasterDataMiddleware()) // Apply middleware
	{
		protected.GET("/profile", userHandler.GetUserProfile)
		//protected.PUT("/profile", userHandler.UpdateUserProfile)
		//protected.POST("/logout", userHandler.Logout)
	}
}
