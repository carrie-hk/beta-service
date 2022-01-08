package handlers

import (
	"beta_service/db"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

type Handler struct {
	database *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{
		database: db,
	}
}

func (h *Handler) HandleGetAssets(w http.ResponseWriter, r *http.Request) {

	log.Print("Directed to Handle Get Assets")
	store := db.NewAssetStore(h.database)
	users, err := store.Assets()
	if err != nil {
		fmt.Printf("error %w", err)
	}

	print(users)

}

// func (h *Handler) HandleGetFeaturedAssets(w http.ResponseWriter r *http.Request){

// 	vars := http.Vars(r)

// }
