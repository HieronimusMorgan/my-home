package master

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ProductID     uint      `gorm:"primaryKey;autoIncrement" json:"product_id"`
	Name          string    `gorm:"not null" json:"name"`
	Description   string    `json:"description,omitempty"`
	Price         float64   `gorm:"not null" json:"price"`
	StockQuantity int       `gorm:"not null" json:"stock_quantity"`
	CategoryID    uint      `gorm:"not null" json:"category_id"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	CreatedBy     string    `json:"created_by,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
	UpdatedBy     string    `json:"updated_by,omitempty"`
	DeletedAt     bool      `json:"deleted_at,omitempty"`
	DeletedBy     string    `json:"deleted_by,omitempty"`
}

var ProductDB *gorm.DB
