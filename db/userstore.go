package db

import (
	"beta_service/models"
	"database/sql"
	"fmt"
)

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{
		DB: db,
	}
}

type UserStore struct {
	*sql.DB
}

func (s *UserStore) User(Email string) (models.User, error) {
	u, err := s.Query("SELECT * from customers WHERE email = $1", Email)
	if err != nil {
		return models.User{}, fmt.Errorf("error returning asset: %w", err)
	}

	print(u)

	//convert sql.rows to models.User
	return models.User{}, nil
}

func (s *UserStore) Users() ([]models.User, error) {
	// var uu []models.User
	// err := s.Query(&uu, "SELECT * from customers")
	// if err != nil {
	// 	return []models.User{}, fmt.Errorf("error returning asset: %w", err)
	// }
	return []models.User{}, nil
}

func (s *UserStore) CreateUser(u *models.User) error {
	panic("not implemented") // TODO: Implement
}
