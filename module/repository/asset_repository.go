package repository

import (
	"Master_Data/module/domain/master"
	"errors"
	"gorm.io/gorm"
)

type AssetRepository struct {
	DB *gorm.DB
}

func NewAssetRepository(db *gorm.DB) *AssetRepository {
	return &AssetRepository{DB: db}
}

func (r AssetRepository) AddAssetCategory(assetCategory *master.AssetCategory) error {
	err := r.DB.Where("name = ?", assetCategory.Name).First(&assetCategory).Error
	if err == nil {
		return errors.New("Asset category already exist")
	}
	return r.DB.Create(assetCategory).Error
}

func (r AssetRepository) GetAssetCategoryByName(name *string) (master.AssetCategory, error) {
	var assetCategory master.AssetCategory
	err := r.DB.Where("name = ?", name).First(&assetCategory).Error
	return assetCategory, err
}

func (r AssetRepository) GetAssetCategoryById(id *string) (interface{}, error) {
	var assetCategory master.AssetCategory
	err := r.DB.Where("asset_category_id = ?", id).First(&assetCategory).Error
	return assetCategory, err
}
