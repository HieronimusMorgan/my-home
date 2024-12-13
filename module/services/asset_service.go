package services

import (
	"Master_Data/module/domain/master"
	"Master_Data/module/repository"
	"Master_Data/utils"
	"gorm.io/gorm"
)

type AssetService struct {
	AssetRepository *repository.AssetRepository
	UserRepository  *repository.UserRepository
}

func NewAssetService(db *gorm.DB) *AssetService {
	assetRepo := repository.NewAssetRepository(db)
	userRepo := repository.NewUserRepository(db)
	return &AssetService{AssetRepository: assetRepo, UserRepository: userRepo}
}

func (s AssetService) AddAssetCategory(name *string, description *string, clientID string) (interface{}, error) {
	var user interface{}
	utils.GetDataFromRedis("user", clientID, &user)
	assetCategory := master.AssetCategory{
		Name:        *name,
		Description: *description,
		CreatedBy:   user.(master.User).FullName,
		UpdatedBy:   user.(master.User).FullName,
	}
	err := s.AssetRepository.AddAssetCategory(&assetCategory)
	if err != nil {
		return nil, err
	}
	a, err := s.AssetRepository.GetAssetCategoryByName(name)
	return a, nil
}

func (s AssetService) GetAssetCategoryByName(name *string) (interface{}, error) {
	assetCategory, err := s.AssetRepository.GetAssetCategoryByName(name)
	if err != nil {
		return nil, err
	}
	return assetCategory, nil
}

func (s AssetService) GetAssetCategoryById(assetID *string) (interface{}, error) {
	assetCategory, err := s.AssetRepository.GetAssetCategoryById(assetID)
	if err != nil {
		return nil, err
	}
	return assetCategory, nil
}
