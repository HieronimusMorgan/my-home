package domain

import (
	"gorm.io/gorm"
	"time"
)

type Token struct {
	TokenID      uint      `gorm:"primaryKey" json:"token_id"`
	UserID       uint      `gorm:"not null" json:"user_id"`
	Token        string    `gorm:"not null" json:"token"`
	RefreshToken string    `gorm:"not null" json:"refresh_token"`
	Expired      bool      `gorm:"not null" json:"expired"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	CreatedBy    string    `json:"created_by,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	UpdatedBy    string    `json:"updated_by,omitempty"`
	DeletedAt    time.Time `json:"deleted_at,omitempty"`
	DeletedBy    string    `json:"deleted_by,omitempty"`
}

var TokenDB *gorm.DB
