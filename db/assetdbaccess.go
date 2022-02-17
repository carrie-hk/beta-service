package db

import (
	"beta_service/models"
	"log"

	"github.com/jmoiron/sqlx"
)

type AssetDbAccess struct {
	*sqlx.DB
}

func (s *AssetDbAccess) AllAssets(index int, pageSize int) ([]models.Asset, error) {
	var aa []models.Asset
	query := "SELECT * from AXU.whisky_bottles WHERE id < $1 ORDER BY id DESC LIMIT $2"

	err := s.Select(&aa, query, index)
	if err != nil {
		log.Println("error selecting asset", err)
		return []models.Asset{}, err
	}

	return aa, nil
}

func (s *AssetDbAccess) FeaturedAssets() ([]models.Asset, error) {
	var fa []models.Asset
	err := s.Select(&fa, "SELECT * from AXU.whisky_bottles WHERE `Featured`=1")
	if err != nil {
		log.Println("error selecting featured asset", err)
		return []models.Asset{}, err
	}
	return fa, nil
}

func (s *AssetDbAccess) TestAssets() ([]models.Asset, error) {
	var fa []models.Asset
	err := s.Select(&fa, "SELECT Age, from AXU.whisky_bottles WHERE `Bottle ID`=1")
	if err != nil {
		log.Println("error selecting featured asset", err)
		return []models.Asset{}, err
	}
	return fa, nil
}
