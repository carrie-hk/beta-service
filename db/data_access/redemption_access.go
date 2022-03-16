package data_access

import (
	"beta_service/api/models"
	"log"
)

func (db *DbAccess) CreateKYC(kyc models.KYC) error {

	_, err := db.NamedExec(`INSERT INTO baxusnft.kyc (wallet_pk, first_name, last_name, phone_num, email, ship_addr_a, 
		ship_addr_b, ship_city, ship_state, ship_zip, dob_day, dob_month, dob_year, title) 
		VALUES (:wallet_pk, :first_name, :last_name, :phone_num, :email, :ship_addr_a, :ship_addr_b, 
			:ship_city, :ship_state, :ship_zip, :dob_day, :dob_month, :dob_year, :title)`,
		kyc)

	if err != nil {
		log.Println("Error inserting KYC information:", err)
		return err
	}

	return nil

}

func (db *DbAccess) GetRedemptionAssets(mintAddr string) ([]models.AssetView, error) {
	var aa []models.AssetView

	query := "SELECT * from baxusnft.asset_view_table WHERE `mint_addr` = ?"

	err := db.Select(&aa, query, mintAddr)
	if err != nil {
		log.Println("Error selecting asset:", err)
		return []models.AssetView{}, err
	}

	return aa, nil
}
