package master

import (
	"gorm.io/gorm"
	"time"
)

type PasswordManager struct {
	PasswordID  uint      `gorm:"primaryKey" json:"password_id,omitempty"`
	UserID      uint      `gorm:"not null" json:"user_id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Password    string    `json:"password,omitempty"`
	Description string    `json:"description,omitempty"`
	Expired     bool      `json:"expired,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	CreatedBy   string    `json:"created_by,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	UpdatedBy   string    `json:"updated_by,omitempty"`
	DeletedAt   bool      `json:"deleted_at,omitempty"`
	DeletedBy   string    `json:"deleted_by,omitempty"`
}

var PasswordManagerDB *gorm.DB
