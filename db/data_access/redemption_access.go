package data_access

import (
	"beta_service/api/models"
	"log"

	"github.com/jmoiron/sqlx"
)

func (db *DbAccess) CreateKYC(kyc models.KYC) error {

	_, err := db.NamedExec(`INSERT INTO kyc.customers (firstName, lastName, phoneNumber, email, streetAddrA_ship, 
		streetAddrB_ship, city_ship, state_ship, country_ship, zipcode_ship, streetAddrA_bill, 
		streetAddrB_bill, city_bill, state_bill, country_bill, zipcode_bill, birthDay, birthMonth, birthYear, title) 
		VALUES (:firstName, :lastName, :phoneNumber, :email, :streetAddrA_ship, 
			:streetAddrB_ship, :city_ship, :state_ship, :country_ship, :zipcode_ship, :streetAddrA_bill, 
			:streetAddrB_bill, :city_bill, :state_bill, :country_bill, :zipcode_bill, :birthDay, :birthMonth, :birthYear, :title )`,
		kyc)

	if err != nil {
		log.Println("Error inserting KYC information:", err)
		return err
	}

	return nil

}

func (db *DbAccess) GetRedemptionAssets(rr_list []models.RedemptionRequest) ([]models.AssetView, error) {
	var aa []models.AssetView
	var rr_string_list []string

	for _, rr_item := range rr_list {
		rr_string_list = append(rr_string_list, rr_item.Mint_Addr)
	}

	log.Println(rr_string_list)

	query := "SELECT * from asset_view_table WHERE mint_addr IN (?);"
	query, args, err := sqlx.In(query, rr_string_list)
	if err != nil {
		return nil, err
	}

	query = db.Rebind(query)
	res, err := db.Queryx(query, args...)
	if err != nil {
		log.Println("Error selecting asset:", err)
		return []models.AssetView{}, err
	}

	for res.Next() {
		var aa_elem models.AssetView
		err = res.StructScan(&aa_elem)
		if err != nil {
			log.Println(err)
		}
		aa = append(aa, aa_elem)
		log.Println(aa_elem)
	}

	return aa, nil
}
