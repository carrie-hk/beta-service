package web

import (
	"beta_service/db"

	"github.com/gorilla/mux"
)

type Handler struct {
	Router *mux.Router
	*AssetHandler
	*UserHandler
}

// 	This function takes the model stores as inputs and returns the router and subrouters for all models
func NewHandler(store *db.Store) (*Handler, error) {

	router := mux.NewRouter()

	return &Handler{
		Router:       router,
		AssetHandler: NewAssetHandler(store, router),
		UserHandler:  NewUserHandler(store, router),
	}, nil
}
