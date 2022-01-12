package main

import (
	"beta_service/web"

	"github.com/gorilla/mux"
)

func Routers(handlers *web.Handler) (*mux.Router, error) {

	router := mux.NewRouter()

	//creation of the SubRouters for the asset model
	subAssetRouter := router.PathPrefix("/").Subrouter()
	subAssetRouter.HandleFunc("/assets", handlers.HandleGetAssets())
	subAssetRouter.HandleFunc("/", handlers.HandleGetFeaturedAssets())

	subRedemptionRouter := router.PathPrefix("/redeem").Subrouter()
	subRedemptionRouter.HandleFunc("/kycform", handlers.HandleCreateUser())

	return router, nil
}
