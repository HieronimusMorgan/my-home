package routes

import (
	"Master_Data/module/handler"
	"Master_Data/module/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PasswordManagerRoutes(r *gin.Engine, db *gorm.DB) {
	passwordHandler := handler.NewPasswordManagerHandler(db)

	protected := r.Group("/api/v1/password")
	protected.Use(middleware.MasterDataMiddleware()) // Apply middleware
	{
		protected.POST("/add", passwordHandler.AddPassword)
		//protected.POST("/update", passwordHandler.UpdatePassword)
		//protected.POST("/delete", passwordHandler.DeletePassword)
		protected.GET("/get", passwordHandler.GetPassword)
	}
}
