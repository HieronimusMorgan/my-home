package master

import (
	"gorm.io/gorm"
	"time"
)

type ProductCategory struct {
	CategoryID  uint      `gorm:"primaryKey;autoIncrement" json:"category_id"`
	Name        string    `gorm:"not null;unique" json:"name"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	CreatedBy   string    `json:"created_by,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	UpdatedBy   string    `json:"updated_by,omitempty"`
	DeletedAt   bool      `json:"deleted_at,omitempty"`
	DeletedBy   string    `json:"deleted_by,omitempty"`
}

var ProductCategoryDB *gorm.DB
