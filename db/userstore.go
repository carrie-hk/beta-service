package db

import (
	"beta_service/models"
	"log"

	"github.com/jmoiron/sqlx"
)

func NewUserStore(db *sqlx.DB) *UserStore {
	return &UserStore{
		DB: db,
	}
}

type UserStore struct {
	*sqlx.DB
}

// func (s *UserStore) User(Email string) (models.User, error) {
// 	u, err := s.Query("SELECT * from customers WHERE email = $1", Email)
// 	if err != nil {
// 		return models.User{}, fmt.Errorf("error returning asset: %w", err)
// 	}

// 	print(u)

// 	//convert sql.rows to models.User
// 	return models.User{}, nil
// }

// func (s *UserStore) Users() ([]models.User, error) {
// 	// var uu []models.User
// 	// err := s.Query(&uu, "SELECT * from customers")
// 	// if err != nil {
// 	// 	return []models.User{}, fmt.Errorf("error returning asset: %w", err)
// 	// }
// 	return []models.User{}, nil
// }

func (s *UserStore) CreateUser(u models.User) bool {

	_, err := s.NamedExec(`INSERT INTO KYC.customers (firstName,lastName, phoneNumber, streetAddrA, 
		streetAddrB, city, email, state, country, zipcode, birthDay, birthMonth, birthYear, title) 
		VALUES (:firstName, :lastName, :phoneNumber, :streetAddrA, :streetAddrB, :city, :email,
		:state, :country, :zipcode, :birthDay, :birthMonth, :birthYear, :title )`,
		u)

	if err != nil {
		log.Println("Issue with insert", err)
		return false
	}

	return true

}
