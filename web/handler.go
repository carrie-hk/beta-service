package web

import (
	"beta_service/db"
)

type Handler struct {
	*AssetHandler
	*UserHandler
}

// 	This function takes the model stores as inputs and returns the router and subrouters for all models
func NewHandler(store *db.Store) (*Handler, error) {
	return &Handler{
		AssetHandler: &AssetHandler{store: store},
		UserHandler:  &UserHandler{store: store},
	}, nil
}
