package data_access

import (
	"beta_service/db/models"
	"log"
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
