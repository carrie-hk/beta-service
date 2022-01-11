package db

import (
	"beta_service/models"
	"log"

	"github.com/jmoiron/sqlx"
)

func NewAssetStore(db *sqlx.DB) *AssetStore {
	return &AssetStore{
		DB: db,
	}
}

type AssetStore struct {
	*sqlx.DB
}

func (s *AssetStore) Assets() []models.Asset {
	var aa []models.Asset
	err := s.Select(&aa, "SELECT * from AXU.whisky_bottles")
	if err != nil {
		log.Println("error selecting asset", err)
		return []models.Asset{}
	}
	return aa
}

func (s *AssetStore) FeaturedAssets() []models.Asset {
	var fa []models.Asset
	err := s.Select(&fa, "SELECT * from AXU.whisky_bottles WHERE `Bottle ID`=2")
	if err != nil {
		log.Println("error selecting featured asset", err)
		return []models.Asset{}
	}
	return fa
}
