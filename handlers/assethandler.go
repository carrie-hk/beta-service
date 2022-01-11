package handlers

import (
	"beta_service/db"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type AssetHandler struct {
	database *sqlx.DB
}

func NewAssetHandler(db *sqlx.DB) *AssetHandler {
	return &AssetHandler{
		database: db,
	}
}

func (h *AssetHandler) HandleGetAssets(w http.ResponseWriter, r *http.Request) {
	log.Print("Directed to Handle Get Assets")
	store := db.NewAssetStore(h.database)
	assets := store.Assets()
	b, err := json.Marshal(assets)
	if err != nil {
		log.Println("Error json marshalling")
	}
	fmt.Fprintf(w, "%s", b)
	log.Print("Returning All Assets")
}

func (h *AssetHandler) HandleGetFeaturedAssets(w http.ResponseWriter, r *http.Request) {
	log.Print("Directed to Handle Get Featured Assets")
	store := db.NewAssetStore(h.database)
	assets := store.FeaturedAssets()
	b, err := json.Marshal(assets)
	if err != nil {
		log.Println("Error json marshalling")
	}
	fmt.Fprintf(w, "%s", b)
	log.Print("Returning All Assets")

}
