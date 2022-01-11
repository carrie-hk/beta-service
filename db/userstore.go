package db

import (
	"beta_service/models"
	"log"

	"github.com/jmoiron/sqlx"
)

type UserStore struct {
	*sqlx.DB
}

func (s *UserStore) CreateUser(u models.User) error {

	_, err := s.NamedExec(`INSERT INTO KYC.customers (firstName,lastName, phoneNumber, streetAddrA, 
		streetAddrB, city, email, state, country, zipcode, birthDay, birthMonth, birthYear, title) 
		VALUES (:firstName, :lastName, :phoneNumber, :streetAddrA, :streetAddrB, :city, :email,
		:state, :country, :zipcode, :birthDay, :birthMonth, :birthYear, :title )`,
		u)

	if err != nil {
		log.Println("Issue with insert", err)
		return err
	}

	return nil

}
