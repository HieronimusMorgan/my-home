package database

import (
	"Master_Data/config"
	"Master_Data/module/domain/master"
	"github.com/golang-migrate/migrate/v4"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var db *gorm.DB

func InitDB() *gorm.DB {
	cfg := config.LoadConfig()
	dsn := "host=" + cfg.DBHost + " user=" + cfg.DBUser + " password=" + cfg.DBPassword +
		" dbname=" + cfg.DBName + " port=" + cfg.DBPort + " sslmode=disable"

	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:         logger.Default.LogMode(logger.Info),
		NamingStrategy: schemaNamingStrategy(cfg.DBSchema), // Set the schema
	})

	db.AutoMigrate(&master.User{}, &master.Balance{},
		&master.Asset{}, &master.AssetCategory{},
		&master.AssetMaintenance{},
		&master.Product{}, &master.ProductCategory{},
		&master.Roles{}, &master.Token{}, &master.PasswordManager{})
	//Run migrations
	//migrationsPath := "file://./migrations"
	//RunMigrations(dsn, migrationsPath)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connection established")
	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	sqlDB.Close()
}

func schemaNamingStrategy(schemaName string) schema.NamingStrategy {
	return schema.NamingStrategy{
		TablePrefix: schemaName + ".", // Use the schema as a prefix
	}
}

func RunMigrations(dsn string, migrationsPath string) {
	m, err := migrate.New(migrationsPath, dsn)
	if err != nil {
		log.Fatalf("Failed to initialize migrations: %v", err)
	}

	// Run migrations up
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Migrations applied successfully")
}
