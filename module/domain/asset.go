package domain

import (
	"gorm.io/gorm"
	"time"
)

type Asset struct {
	AssetID         uint      `gorm:"primaryKey;autoIncrement" json:"asset_id"`
	Name            string    `gorm:"not null" json:"name"`
	Description     string    `json:"description,omitempty"`
	Value           float64   `gorm:"not null" json:"value"`
	AcquisitionDate time.Time `json:"acquisition_date,omitempty"`
	AssetCategoryID uint      `gorm:"not null" json:"asset_category_id"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	CreatedBy       string    `json:"created_by,omitempty"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
	UpdatedBy       string    `json:"updated_by,omitempty"`
	DeletedAt       time.Time `json:"deleted_at,omitempty"`
	DeletedBy       string    `json:"deleted_by,omitempty"`
}

var AssetDB *gorm.DB
