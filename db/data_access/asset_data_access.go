package data_access

import (
	"beta_service/api/models"
	"log"
)

func (db *DbAccess) GetAllAssets(pageIndex int, pageSize int) ([]models.Asset, error) {
	var aa []models.Asset
	query := "SELECT * from baxusnft.asset_view_table WHERE `axu_id` > ? ORDER BY 'axu_id' DESC LIMIT ?"

	err := db.Select(&aa, query, pageIndex, pageSize)
	if err != nil {
		log.Println("Error selecting all assets:", err)
		return []models.Asset{}, err
	}

	return aa, nil
}

func (db *DbAccess) GetFeaturedAssets() ([]models.Asset, error) {
	var fa []models.Asset
	query := "SELECT * from baxusnft.asset_view_table WHERE `featured`=1"

	err := db.Select(&fa, query)
	if err != nil {
		log.Println("Error selecting featured assets:", err)
		return []models.Asset{}, err
	}
	return fa, nil
}
