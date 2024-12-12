package domain

import (
	"gorm.io/gorm"
	"time"
)

type AssetMaintenance struct {
	AssetMaintenanceID  uint      `gorm:"primaryKey;autoIncrement" json:"asset_maintenance_id"`
	AssetID             uint      `gorm:"not null" json:"asset_id"`
	Cost                float64   `gorm:"not null" json:"cost"`
	Notes               string    `json:"notes,omitempty"`
	MaintenanceDate     time.Time `json:"maintenance_date,omitempty"`
	NextMaintenanceDate time.Time `json:"next_maintenance_date,omitempty"`
	CreatedAt           time.Time `json:"created_at,omitempty"`
	CreatedBy           string    `json:"created_by,omitempty"`
	UpdatedAt           time.Time `json:"updated_at,omitempty"`
	UpdatedBy           string    `json:"updated_by,omitempty"`
	DeletedAt           time.Time `json:"deleted_at,omitempty"`
	DeletedBy           string    `json:"deleted_by,omitempty"`
}

var AssetMaintenanceDB *gorm.DB
