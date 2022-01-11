package handlers

import (
	"beta_service/db"
	"beta_service/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/jmoiron/sqlx"
)

type Handler struct {
	database *sqlx.DB
}

func NewHandler(db *sqlx.DB) *Handler {
	return &Handler{
		database: db,
	}
}

func (h *Handler) HandleGetAssets(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) HandleGetFeaturedAssets(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	log.Print("Directed to Handle Create User")
	store := db.NewUserStore(h.database)
	err := r.ParseForm()
	if err != nil {
		log.Println(w, "ParseForm() err: %v", err)
	}
	decoder := schema.NewDecoder()
	var user models.User
	err = decoder.Decode(user, r.PostForm)
	if err != nil {
		log.Println(w, "DecodeForm err: %v", err)
	}
	success := store.CreateUser(user)
	fmt.Fprintf(w, "%v", success)

}
