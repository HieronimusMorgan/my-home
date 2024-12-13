package routes

import (
	"Master_Data/module/handler"
	"Master_Data/module/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func AssetRoutes(r *gin.Engine, db *gorm.DB) {
	authHandler := handler.NewAssetHandler(db)

	protected := r.Group("/api/v1/asset")
	protected.Use(middleware.MasterDataMiddleware())
	{
		protected.POST("/add/category", authHandler.AddAssetCategory)
		protected.GET("/category/:id", func(c *gin.Context) {
			id := c.Param("id")
			if _, err := strconv.Atoi(id); err == nil {
				authHandler.GetAssetCategoryById(c)
			} else {
				authHandler.GetAssetCategoryByName(c)
			}
		})
		//public.POST("/login", authHandler.Login)
		//public.POST("/refresh", authHandler.RefreshToken)
		//public.GET("/logout", authHandler.Logout)
	}
}
