package data_access

import (
	"beta_service/api/models"
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

	// Perform a query to update the asset_status field in the axu table
	query_axu := "UPDATE axu SET asset_status = :asset_status WHERE axu_id = :axu_id AND mint_addr = :mint_addr"

	res_axu, err := db.NamedExec(query_axu, su_item)
	if err != nil {
		log.Println("Error updating asset status:", err)
		return err
	}

	// Check to ensure that rows were actually affected by update query
	_, err = res_axu.RowsAffected()
	if err != nil {
		log.Println("Could not update axu table - invalid AXU ID or Mint Address")
		return err
	}

	// Perform the same query to update the asset_view_table
	query_avt := "UPDATE asset_view_table SET asset_status = :asset_status WHERE axu_id = :axu_id AND mint_addr = :mint_addr"

	res_avt, err := db.NamedExec(query_avt, su_item)
	if err != nil {
		log.Println("Error updating asset status:", err)
		return err
	}

	// Check to ensure that rows were actually affected by update query
	_, err = res_avt.RowsAffected()
	if err != nil {
		log.Println("Could not update asset view table - invali AXU ID or Mint Address")
		return err
	}

	return nil
}
