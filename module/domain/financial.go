package domain

import (
	"gorm.io/gorm"
	"time"
)

type Financial struct {
	FinancialID uint      `gorm:"primaryKey" json:"financial_id"`
	UserID      uint      `gorm:"not null, unique" json:"user_id"`
	Balance     float64   `gorm:"not null" json:"balance"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	CreatedBy   string    `json:"created_by,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	UpdatedBy   string    `json:"updated_by,omitempty"`
	DeletedAt   time.Time `json:"deleted_at,omitempty"`
	DeletedBy   string    `json:"deleted_by,omitempty"`
}

var FinanceDB *gorm.DB
