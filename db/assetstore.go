package db

import (
	"beta_service/models"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

func NewAssetStore(db *sql.DB) *AssetStore {
	return &AssetStore{
		DB: db,
	}
}

type AssetStore struct {
	*sql.DB
}

func (s *AssetStore) Asset(id uuid.UUID) (models.Asset, error) {
	// var a models.Asset
	a, err := s.Query("SELECT * from AXU.whisky_bottles WHERE id = $1", id)
	if err != nil {
		return models.Asset{}, fmt.Errorf("error returning asset: %w", err)
	}

	print(a)
	return models.Asset{}, nil
}

func (s *AssetStore) Assets() ([]models.Asset, error) {
	// var aa []models.Asset
	aa, err := s.Query("SELECT * from AXU.whisky_bottles")
	if err != nil {
		return nil, fmt.Errorf("error returning asset: %w", err)
	}
	print(aa)
	return []models.Asset{}, nil
}

func (s *AssetStore) FeaturedAssets(Featured bool) ([]models.Asset, error) {
	// var fa []models.Asset
	fa, err := s.Query("SELECT * from AXU.whisky_bottles WHERE featured = $1", Featured)
	if err != nil {
		return []models.Asset{}, fmt.Errorf("error returning asset: %w", err)
	}
	print(fa)
	return []models.Asset{}, nil
}
