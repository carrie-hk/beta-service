package db

import (
	"beta_service/models"
	"fmt"

	"github.com/google/uuid"
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

func (s *AssetStore) Asset(id uuid.UUID) (models.Asset, error) {
	var a models.Asset
	err := s.Get(&a, "SELECT * from assets WHERE id = $1", id)
	if err != nil {
		return models.Asset{}, fmt.Errorf("error returning asset: %w", err)
	}
	return a, nil
}

func (s *AssetStore) Assets() ([]models.Asset, error) {
	var aa []models.Asset
	err := s.Select(&aa, "SELECT * from assets")
	if err != nil {
		return []models.Asset{}, fmt.Errorf("error returning asset: %w", err)
	}
	return aa, nil
}

func (s *AssetStore) FeaturedAssets(Featured bool) ([]models.Asset, error) {
	var fa []models.Asset
	err := s.Get(&fa, "SELECT * from assets WHERE featured = $1", Featured)
	if err != nil {
		return []models.Asset{}, fmt.Errorf("error returning asset: %w", err)
	}
	return fa, nil
}
