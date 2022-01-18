package db

import (
	"beta_service/models"
	"log"

	"github.com/jmoiron/sqlx"
)

type UserDbAccess struct {
	*sqlx.DB
}

func (s *UserDbAccess) CreateUser(u models.User) error {

	_, err := s.NamedExec(`INSERT INTO KYC.customers (firstName, lastName, phoneNumber, email, streetAddrA_ship, 
		streetAddrB_ship, city_ship, state_ship, country_ship, zipcode_ship, streetAddrA_bill, 
		streetAddrB_bill, city_bill, state_bill, country_bill, zipcode_bill, birthDay, birthMonth, birthYear, title) 
		VALUES (:firstName, :lastName, :phoneNumber, :email, :streetAddrA_ship, 
			:streetAddrB_ship, :city_ship, :state_ship, :country_ship, :zipcode_ship, :streetAddrA_bill, 
			:streetAddrB_bill, :city_bill, :state_bill, :country_bill, :zipcode_bill, :birthDay, :birthMonth, :birthYear, :title )`,
		u)

	if err != nil {
		log.Println("Issue with insert", err)
		return err
	}

	return nil

}
