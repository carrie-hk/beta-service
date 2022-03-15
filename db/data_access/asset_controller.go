package data_access

import (
	"beta_service/api/models"
	"log"
)

func (db *DbAccess) GetAllAssets(pageIndex int, pageSize int) ([]models.Asset, error) {
	var aa []models.Asset
	query := "SELECT * from baxusnft.axu WHERE `axu_id` > ? ORDER BY 'axu_id' DESC LIMIT ?"

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

	query := "UPDATE baxusnft.axu SET asset_status = :asset_status WHERE axu_id = :axu_id AND mint_addr = :mint_addr"

	for _, su_item := range su_list {

		log.Println(su_item)
		res, err := db.NamedExec(query, su_item)
		if err != nil {
			log.Println("Error updating asset status:", err)
			return err
		}
		log.Println("Query result: ", res)
	}

	return nil
}
