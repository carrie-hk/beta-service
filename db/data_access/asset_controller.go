package data_access

import (
	"beta_service/api/models"
	"errors"
	"log"
)

func (db *DbAccess) GetAllAssets(pageIndex int, pageSize int) ([]models.AssetView, error) {
	var aa []models.AssetView
	query := "SELECT * from asset_view_table WHERE `axu_id` > ? ORDER BY 'axu_id' DESC LIMIT ?"

	err := db.Select(&aa, query, pageIndex, pageSize)
	if err != nil {
		log.Println("Error selecting all assets:", err)
		return []models.AssetView{}, err
	}

	return aa, nil
}

func (db *DbAccess) GetFeaturedAssets() ([]models.AssetView, error) {
	var fa []models.AssetView
	query := "SELECT * from asset_view_table WHERE `featured`=1"

	err := db.Select(&fa, query)
	if err != nil {
		log.Println("Error selecting featured assets:", err)
		return []models.AssetView{}, err
	}
	return fa, nil
}

func (db *DbAccess) UpdateAssetStatus(su_item models.StatusUpdate) error {

	query := "UPDATE axu SET asset_status = :new_status WHERE axu_id = :axu_id AND mint_addr = :mint_addr"

	res, err := db.NamedExec(query, su_item)
	if err != nil {
		log.Println("Error updating asset status:", err)
		return err
	}

	// Check to ensure that rows were actually affected by update query
	row_aff, _ := res.RowsAffected()
	if row_aff == 0 {
		log.Println("Invalid AXU ID or Mint Address")
		return errors.New("Invalid AXU ID or Mint Address")
	}

	return nil
}
