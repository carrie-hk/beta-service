package models

import "github.com/google/uuid"

type Asset struct {
	ID       uuid.UUID `db:"id"`
	Featured bool      `db:"featured"`
}

type AssetStore interface {
	Asset(id uuid.UUID) (Asset, error)
	Assets() ([]Asset, error)
	FeaturedAssets(Featured bool) ([]Asset, error)
}
