package db

import (
	"beta_service/models"
	"fmt"

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

func (s *UserStore) User(Email string) (models.User, error) {
	var u models.User
	err := s.Get(&u, "SELECT * from customers WHERE email = $1", Email)
	if err != nil {
		return models.User{}, fmt.Errorf("error returning asset: %w", err)
	}
	return u, nil
}

func (s *UserStore) Users() ([]models.User, error) {
	var uu []models.User
	err := s.Select(&uu, "SELECT * from customers")
	if err != nil {
		return []models.User{}, fmt.Errorf("error returning asset: %w", err)
	}
	return uu, nil
}

func (s *UserStore) CreateUser(u *models.User) error {
	panic("not implemented") // TODO: Implement
}
