package main

import (
	"Master_Data/module/routes"
	"Master_Data/package/database"
	"Master_Data/utils"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/schema"
)

func main() {
	// Initialize Redis
	utils.InitializeRedis()

	// Initialize database
	db := database.InitDB()
	defer database.CloseDB(db)

	// Setup Gin router
	r := gin.Default()

	// Register routes
	routes.UserRoutes(r, db)
	routes.AuthRoutes(r, db)

	// Run server
	log.Println("Starting server on :8080")
	r.Run(":8080")
}

func schemaNamingStrategy(schemaName string) schema.NamingStrategy {
	return schema.NamingStrategy{
		TablePrefix: schemaName + ".", // Use the schema as a prefix
	}
}
