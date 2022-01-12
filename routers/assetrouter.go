package routers

import (
	"beta_service/handlers"

	"github.com/gorilla/mux"
)

func NewAssetRouter(router *mux.Router, assetHandler *handlers.AssetHandler) (*mux.Router, error) {

	// Create Subrouter for the asset model and assign function handlers for paths
	assetRouter := router.PathPrefix("/assets").Subrouter()

	assetRouter.HandleFunc("/all", assetHandler.HandleGetAllAssets())
	assetRouter.HandleFunc("/featured", assetHandler.HandleGetFeaturedAssets())

	return router, nil
}

// (*mux.Router, error)
