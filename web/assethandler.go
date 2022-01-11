package web

import (
	"beta_service/db"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type AssetHandler struct {
	router *mux.Router
	store  *db.Store
}

func NewAssetHandler(store *db.Store, router *mux.Router) *AssetHandler {
	h := &AssetHandler{
		router: router,
		store:  store,
	}

	//creation of the SubRouters for the asset model
	subrouter := router.PathPrefix("/").Subrouter()
	subrouter.HandleFunc("/assets", h.HandleGetAssets())
	subrouter.HandleFunc("/", h.HandleGetFeaturedAssets())

	return h
}

func (h *AssetHandler) HandleTestGetAssets() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("Directed to HandleTestGetAssets()")

		assets, err := h.store.AssetStore.TestAssets()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println("Error getting assets from database")
			return
		}
		b, err := json.Marshal(assets)
		if err != nil {
			log.Println(err)
		}

		fmt.Fprintf(w, "%s", b)
	}

}

//This function returns all of the assets in the AXU.whisky_bottles
func (h *AssetHandler) HandleGetAssets() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("Directed to HandleGetAssets()")

		assets, err := h.store.AssetStore.Assets()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err)
			return
		}

		b, err := json.Marshal(assets)
		if err != nil {
			log.Println(err)
		}

		fmt.Fprintf(w, "%s", b)
	}
}

//This function returns a featured subset of the bottles
func (h *AssetHandler) HandleGetFeaturedAssets() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("Directed to HandleGetFeaturedAssets()")

		assets, err := h.store.AssetStore.FeaturedAssets()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err)
			return
		}

		b, err := json.Marshal(assets)
		if err != nil {
			log.Println("Error json marshalling")
		}

		fmt.Fprintf(w, "%s", b)
	}
}
