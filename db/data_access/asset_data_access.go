package data_access

import (
	"beta_service/api/models"
	"log"
)

func (db *DbAccess) GetAllAssets(pageIndex int, pageSize int) ([]models.AssetView, error) {
	var aa []models.AssetView
	query := "SELECT * from baxusnft.asset_view_table WHERE `axu_id` > ? ORDER BY 'axu_id' DESC LIMIT ?"

	err := db.Select(&aa, query, pageIndex, pageSize)
	if err != nil {
		log.Println("Error selecting all assets:", err)
		return []models.AssetView{}, err
	}

	return aa, nil
}

func (db *DbAccess) GetFeaturedAssets() ([]models.AssetView, error) {
	var fa []models.AssetView
	query := "SELECT * from baxusnft.asset_view_table WHERE `featured`=1"

	err := db.Select(&fa, query)
	if err != nil {
		log.Println("Error selecting featured assets:", err)
		return []models.AssetView{}, err
	}
	return fa, nil
}
