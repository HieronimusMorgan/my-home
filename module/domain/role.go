package domain

import (
	"gorm.io/gorm"
	"time"
)

type Roles struct {
	RoleID    uint      `gorm:"primaryKey" json:"role_id"`
	Name      string    `gorm:"unique" json:"name"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	CreatedBy string    `json:"created_by,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	UpdatedBy string    `json:"updated_by,omitempty"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	DeletedBy string    `json:"deleted_by,omitempty"`
}

var RolesDB *gorm.DB
