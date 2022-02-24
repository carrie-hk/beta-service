package db_access

import (
	"beta_service/models"
	"log"

	"github.com/jmoiron/sqlx"
)

type AssetDbAccess struct {
	*sqlx.DB
}

func (s *AssetDbAccess) GetAllAssets(pageIndex int, pageSize int) ([]models.Asset, error) {
	var aa []models.Asset
	query := "SELECT * from AXU.whisky_bottles WHERE `Bottle ID` > ? ORDER BY 'Bottle ID' DESC LIMIT ?"

	err := s.Select(&aa, query, pageIndex, pageSize)
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
