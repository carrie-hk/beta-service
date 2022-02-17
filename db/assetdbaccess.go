package db

import (
	"beta_service/models"
	"log"

	"github.com/jmoiron/sqlx"
)

type AssetDbAccess struct {
	*sqlx.DB
}

func (s *AssetDbAccess) GetAllAssets(pageSize int, pageIndex int) ([]models.Asset, error) {
	var aa []models.Asset
	query := "SELECT * from AXU.whisky_bottles WHERE `Bottle ID` < ? ORDER BY 'Bottle ID' DESC LIMIT ?"

	err := s.Select(&aa, query, pageSize, pageIndex)
	if err != nil {
		log.Println("error selecting asset", err)
		return []models.Asset{}, err
	}

	return aa, nil
}

func (s *AssetDbAccess) GetFeaturedAssets() ([]models.Asset, error) {
	var fa []models.Asset
	query := "SELECT * from AXU.whisky_bottles WHERE `Featured`=1"

	err := s.Select(&fa, query)
	if err != nil {
		log.Println("error selecting featured asset", err)
		return []models.Asset{}, err
	}
	return fa, nil
}
