package handlers

import (
	"beta_service/db"
	"beta_service/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/jmoiron/sqlx"
)

type UserHandler struct {
	database *sqlx.DB
}

func NewUserHandler(db *sqlx.DB) *UserHandler {
	return &UserHandler{
		database: db,
	}
}

func (h *UserHandler) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
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
