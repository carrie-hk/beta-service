package db

import "beta_service/models"

type Store struct {
	models.User
	models.Asset
}
