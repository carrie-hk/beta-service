package data_access

import (
	"beta_service/db/models"
	"log"
)

func (db *DbAccess) GetAllAssets(pageIndex int, pageSize int) ([]models.Asset, error) {
	var aa []models.Asset
	query := "SELECT * from AXU.whisky_bottles WHERE `Bottle ID` > ? ORDER BY 'Bottle ID' DESC LIMIT ?"

	err := db.Select(&aa, query, pageIndex, pageSize)
	if err != nil {
		log.Println("Error selecting all assets:", err)
		return []models.Asset{}, err
	}

	return aa, nil
}

func (db *DbAccess) GetFeaturedAssets() ([]models.Asset, error) {
	var fa []models.Asset
	query := "SELECT * from AXU.whisky_bottles WHERE `Featured`=1"

	err := db.Select(&fa, query)
	if err != nil {
		log.Println("Error selecting featured assets:", err)
		return []models.Asset{}, err
	}
	return fa, nil
}

func (db *DbAccess) UpdateAssetStatus(su_list []models.StatusUpdate) error {
	query := "UPDATE axu SET asset_status = ? WHERE axu_id = ?"

	for _, su_item := range su_list {
		err := db.MustExec(query, su_item.New_Status, su_item.AXU_ID)
		if err != nil {
			log.Println("Error updating asset status:", err)
		}
	}

	return nil
}
